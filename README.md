# Particlon

An attempt at a particle system written in Golang, using [Ebiten engine](https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#section-readme).
In hindsight, using a GC language for millions of particles was probably not the best idea.

## Goals

- [x] Particles
- [x] QuadTree
    - [x] Optimize lookup
    - [x] ~~Re-insert particles while they move outside of the quadrant~~
    - [x] Rebuild the tree to generate new nodes
- [x] Attraction/Repulsion between particles
- [ ] Prevent particle collapsing
- [ ] Repulsion/Attraction matrix between particles
- [ ] Runtime config
