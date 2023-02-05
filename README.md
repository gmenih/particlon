# Particlon

An attempt at a particle system written in Golang, using [Ebiten engine](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#section-readme).

## Goals

- [x] Particles
- [x] QuadTree
    - [ ] Optimize lookup
    - [ ] Re-insert particles while they move outside of the quadrant
- [ ] Attraction/Repulsion between particles
- [ ] Prevent particle collapsing
- [ ] Repulsion/Attraction matrix between particles
- [ ] Runtime config
