package example_test

import (
	shineplum "github.com/MagicalLas/shine-plum"
	"github.com/MagicalLas/shine-plum/internal/example"
	"github.com/MagicalLas/shine-plum/internal/example/usecase"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
