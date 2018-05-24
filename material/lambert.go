package material

import (
	"math"
	"math/rand"

	"github.com/hunterloftis/pbr/geom"
	"github.com/hunterloftis/pbr/rgb"
)

type Lambert struct {
	Color rgb.Energy
}

func (l Lambert) Sample(out geom.Direction, rnd *rand.Rand) geom.Direction {
	normal := geom.Up
	return normal.RandHemiCos(rnd)
}

func (l Lambert) PDF(in, out geom.Direction) float64 {
	normal := geom.Up
	return in.Dot(normal) * math.Pi
}

func (l Lambert) Eval(in, out geom.Direction) rgb.Energy {
	normal := geom.Up
	return l.Color.Scaled(math.Pi * in.Dot(normal))
}
