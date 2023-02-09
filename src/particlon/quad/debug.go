package quad

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

var colors = []color.Color{
	colornames.Red,
	colornames.Blue,
	colornames.Pink,
	colornames.Green,
	colornames.Lime,
	colornames.Aquamarine,
	colornames.Chartreuse,
}

func (t *Tree[TElement]) Debug(image *ebiten.Image, depth int) {
	col := colors[depth%len(colors)]
	if !t.isLeaf {
		for _, n := range t.nodes {
			n.Debug(image, depth+1)
		}
	} else {
		for _, e := range t.elements {
			x, y := e.Identity()
			ebitenutil.DrawRect(image, x-3, y-3, 6, 6, col)
		}
	}

	ebitenutil.DrawLine(image, t.boundary.MinX, t.boundary.MinY, t.boundary.MaxX, t.boundary.MinY, col)
	ebitenutil.DrawLine(image, t.boundary.MaxX, t.boundary.MinY, t.boundary.MaxX, t.boundary.MaxY, col)
	ebitenutil.DrawLine(image, t.boundary.MaxX, t.boundary.MaxY, t.boundary.MinX, t.boundary.MaxY, col)
	ebitenutil.DrawLine(image, t.boundary.MinX, t.boundary.MaxY, t.boundary.MinX, t.boundary.MinY, col)
}
