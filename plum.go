package shine_plum

import (
	"encoding/json"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"os"
	"reflect"
	"strings"
)

var errorInterface = reflect.TypeOf((*error)(nil)).Elem()

type Plum struct {
	u  interface{}
	ts reflect.Type
	vs reflect.Value
}

func NewPlum(u interface{}) *Plum {
	t := reflect.TypeOf(u)
	v := reflect.ValueOf(u)
	if t.Kind() == reflect.Ptr {
		return &Plum{u: u, ts: t.Elem(), vs: v}
	} else if t.Kind() == reflect.Struct {
		return &Plum{u: u, ts: t, vs: v}
	}
	return nil
}

func (p *Plum) Name() string {
	return p.ts.Name()
}

func (p *Plum) MethodNames() []string {
	var names []string
	for i := 0; i < p.vs.NumMethod(); i++ {
		names = append(names, p.vs.Type().Method(i).Name)
	}
	return names
}

type MethodDescription struct {
	MethodName        string
	SimpleDescription string
}

func (p *Plum) MethodDescription(name string) MethodDescription {
	fset := token.NewFileSet()
	wd, _ := os.Getwd()
	path := ""
	if len(strings.Split(p.ts.PkgPath(), "/")) > 3 {
		projectName := strings.Split(p.ts.PkgPath(), "/")[2]
		projectPath := strings.Split(wd, projectName)[0]
		packagePath := strings.Split(p.ts.PkgPath(), projectName)[1]
		path = projectPath + projectName + packagePath
	} else {
		path = "./"
	}

	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		fmt.Print(err)
		return MethodDescription{
			"ERROR",
			"ERROR",
		}
	}

	for _, f := range d {
		p := doc.New(f, "./", 0)
		for _, t := range p.Types {
			for _, m := range t.Methods {
				if m.Name == name {
					formattedDoc := strings.ReplaceAll(m.Doc, "\n", " ")
					return MethodDescription{
						m.Name,
						formattedDoc,
					}
				}
			}
		}
	}
	return MethodDescription{
		"ERROR",
		"ERROR ",
	}
}

type Method struct {
	Name       string
	Parameters []interface{}
}

func (m Method) jsonParameter() []byte {
	bs, _ := json.Marshal(m.Parameters)
	return bs
}

type MethodResult struct {
	JSON string
}

func (p *Plum) ExecuteSimpleMethod(method Method) (*MethodResult, error) {
	var params []reflect.Value
	for _, v := range method.Parameters {
		params = append(params, reflect.ValueOf(v))
	}

	rs := p.vs.MethodByName(method.Name).Call(params)
	r := rs[0].Elem().Interface()
	bs, _ := json.Marshal(r)
	return &MethodResult{
		JSON: string(bs),
	}, nil
}

func (p *Plum) ExecuteMethod(method Method) (*MethodResult, error) {
	var params []reflect.Value

	m := p.vs.MethodByName(method.Name)
	mt := m.Type()

	inCount := mt.NumIn()
	for i := 0; i < inCount; i++ {
		it := mt.In(i)
		ni := reflect.New(it)
		_ = json.Unmarshal(method.jsonParameter(), ni.Interface())
		in := reflect.ValueOf(ni.Interface()).Elem()
		params = append(params, in)
	}

	rs := m.Call(params)
	r := rs[0].Interface()
	bs, _ := json.Marshal(r)
	return &MethodResult{
		JSON: string(bs),
	}, nil
}
