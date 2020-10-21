package shine_plum_test

import (
	"encoding/json"
	shineplum "github.com/MagicalLas/shine-plum"
	"github.com/MagicalLas/shine-plum/internal/example"
	"github.com/MagicalLas/shine-plum/internal/example/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlum_Name(t *testing.T) {
	u := &UseCase{}
	plum := shineplum.NewPlum(u)

	assert.Equal(t, "UseCase", plum.Name())
}

func TestPlum_Name_NotPointer(t *testing.T) {
	u := UseCase{}
	plum := shineplum.NewPlum(u)

	assert.Equal(t, "UseCase", plum.Name())
}

func TestPlum_MethodNames(t *testing.T) {
	u := &UseCase{}
	plum := shineplum.NewPlum(u)

	assert.Len(t, plum.MethodNames(), 2)
	assert.Equal(t, "GetAd", plum.MethodNames()[0])
}

func TestPlum_MethodDescription(t *testing.T) {
	u := &UseCase{}
	plum := shineplum.NewPlum(u)

	d := plum.MethodDescription("GetAd")
	assert.Equal(t, "GetAd", d.MethodName)
	assert.Equal(t, "GetAd is get ad. las aka wonho ", d.SimpleDescription)
}

func TestPlum_MethodDescription1(t *testing.T) {
	u := &example.UseCase{}
	plum := shineplum.NewPlum(u)

	d := plum.MethodDescription("GetAd")
	assert.Equal(t, "GetAd", d.MethodName)
	assert.Equal(t, "GetAd is get ad. in example ", d.SimpleDescription)
}

func TestPlum_MethodDescription2(t *testing.T) {
	u := &usecase.UseCase{}
	plum := shineplum.NewPlum(u)

	d := plum.MethodDescription("GetAd")
	assert.Equal(t, "GetAd", d.MethodName)
	assert.Equal(t, "GetAd is get ad. in usecase ", d.SimpleDescription)
}

func TestPlum_ExecuteSimpleMethod(t *testing.T) {
	u := &UseCase{}
	ad := u.GetAd("", 1)
	bs, _ := json.Marshal(ad)
	plum := shineplum.NewPlum(u)

	m := shineplum.Method{
		Name: "GetAd",
		Parameters: []interface{}{
			"", 1,
		},
	}

	result, err := plum.ExecuteSimpleMethod(m)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, string(bs), result.JSON)
}


func TestPlum_ExecuteMethod(t *testing.T) {
	u := &UseCase{}
	ad := u.ListAd(&ListAdCommand{})
	bs, _ := json.Marshal(ad)
	plum := shineplum.NewPlum(u)

	m := shineplum.Method{
		Name: "ListAd",
		Parameters: []interface{}{
			map[string]interface{}{
				"LineitemID": "l",
				"ADNID":      1,
				"UnitID":     2,
				"User": map[string]interface{}{
					"IFA": "asd",
					"IP":  "ww",
				},
			},
		},
	}

	result, err := plum.ExecuteMethod(m)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, string(bs), result.JSON)
}