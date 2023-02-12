package main

import (
	"gmenih/particlon/src/particlon/particlon"
	"log"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	WIDTH  = 1000
	HEIGHT = 1000
)

func main() {

	rand.Seed(time.Now().UnixNano())

	ebiten.SetWindowSize(WIDTH, HEIGHT)
	ebiten.SetWindowTitle("Particlon")

	system := particlon.NewParticleSystem(WIDTH, HEIGHT)

	system.InitParticles()

	// system := particle.NewParticleSystem(particle.SCREEN_WIDTH, particle.SCREEN_HEIGHT)

	if err := ebiten.RunGame(system); err != nil {
		log.Fatal(err)
	}
}
