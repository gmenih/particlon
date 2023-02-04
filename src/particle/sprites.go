package particle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type Sprites = map[color.Color]*ebiten.Image

var colors []color.Color = []color.Color{
	colornames.Red,
	colornames.Orange,
	colornames.Yellow,
	colornames.Green,
	colornames.Blue,
	colornames.Indigo,
	colornames.Violet,
}

func generateSprites() Sprites {
	var sprites Sprites = make(map[color.Color]*ebiten.Image, len(colors))
	size := 6
	half := float64(size) / 2

	for _, c := range colors {
		sprites[c] = ebiten.NewImage(size, size)
		ebitenutil.DrawCircle(sprites[c], half, half, half, c)
	}

	return sprites
}
