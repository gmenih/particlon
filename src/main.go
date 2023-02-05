package main

import (
	"gmenih/particlon/src/particle"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

func main() {

	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Particlon")

	system := particle.NewParticleSystem(screenWidth, screenHeight)

	if err := ebiten.RunGame(system); err != nil {
		log.Fatal(err)
	}
}
