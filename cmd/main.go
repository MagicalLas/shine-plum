package main

import (
	shine_plum "github.com/MagicalLas/shine-plum"
	"github.com/MagicalLas/shine-plum/internal/example"
)

func main() {
	shine := shine_plum.NewShine()
	plum := shine_plum.NewPlum(&example.UseCase{})

	sp := shine_plum.Bright(plum, shine)

	<-sp.TurnOn()
}
