package particle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

var colors = map[ParticleKind]color.Color{
	0: colornames.Red,
	1: colornames.Orange,
	2: colornames.Yellow,
	// 3: colornames.Green,
	// 4: colornames.Khaki,
	// 5: colornames.Violet,
}

func GenerateSprites() []*ebiten.Image {
	sprites := make([]*ebiten.Image, len(colors))
	for i, c := range colors {
		sprites[i] = generateSprite(c)
	}
	return sprites
}

func generateSprite(c color.Color) *ebiten.Image {
	sprite := ebiten.NewImage(6, 6)
	ebitenutil.DrawCircle(sprite, 3, 3, 3, c)

	return sprite
}
