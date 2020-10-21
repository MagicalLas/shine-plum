package shine_plum

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Shine struct {
	port int
}

func (s *Shine) Shadow(e *gin.Engine, p *Plum) {
	names := p.MethodNames()
	for i := range names {
		func() {
			ii := i
			path := fmt.Sprintf("/%s/%s/", p.Name(), names[ii])
			e.GET(path, func(c *gin.Context) {
				c.JSON(200, gin.H{
					p.Name():      names[ii],
					"Description": p.MethodDescription(names[ii]).SimpleDescription,
				})
			})
		}()
	}
}

func NewShine() *Shine {
	return &Shine{}
}

type ShinePlum struct {
	ps []*Plum
	s  *Shine
}

func Bright(p *Plum, s *Shine) *ShinePlum {
	return &ShinePlum{
		ps: []*Plum{p},
		s:  s,
	}
}

func (sp *ShinePlum) Bright(p *Plum) {
	sp.ps = append(sp.ps, p)
}

func (sp *ShinePlum) TurnOn() <-chan interface{}{
	ch := make(chan interface{})
	go func() {
		r := gin.Default()
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		for _, p := range sp.ps {
			sp.s.Shadow(r, p)
		}
		r.Run()
		ch <- nil
	}()
	return ch
}
