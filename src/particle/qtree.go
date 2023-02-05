package particle

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

const MAX_CAPACITY = 16

type QuadTree struct {
	boundary  Bounds
	divided   bool
	particles []*Particle
	nodes     [4]*QuadTree
}

func NewQuadTree(minX, minY, maxX, maxY float64) *QuadTree {
	return &QuadTree{
		boundary: Bounds{
			Min: VV(minX, minY),
			Max: VV(maxX, maxY),
		},
	}
}

func (q *QuadTree) Insert(v *Particle) bool {
	if !q.boundary.Contains(v.Position) {
		return false
	}

	if len(q.particles) < MAX_CAPACITY {
		q.particles = append(q.particles, v)
		return true
	}

	if !q.divided {
		q.subdivide()
	}

	for _, p := range q.particles {
		for _, node := range q.nodes {
			if node.Insert(p) {
				break
			}
		}
	}

	q.particles = nil

	for _, node := range q.nodes {
		if node.Insert(v) {
			return true
		}
	}

	return false
}

func (q *QuadTree) Debug(screen *ebiten.Image, depth int) {
	if q.divided {
		for _, node := range q.nodes {
			node.Debug(screen, depth+1)
		}
	}

	colors := []color.Color{
		colornames.Red,
		colornames.Orange,
		colornames.Yellow,
		colornames.Blue,
		colornames.Violet,
	}

	ebitenutil.DrawLine(screen, q.boundary.Min.X, q.boundary.Min.Y, q.boundary.Max.X, q.boundary.Min.Y, colors[depth%len(colors)])
	ebitenutil.DrawLine(screen, q.boundary.Max.X, q.boundary.Min.Y, q.boundary.Max.X, q.boundary.Max.Y, colors[depth%len(colors)])
	ebitenutil.DrawLine(screen, q.boundary.Max.X, q.boundary.Max.Y, q.boundary.Min.X, q.boundary.Max.Y, colors[depth%len(colors)])
	ebitenutil.DrawLine(screen, q.boundary.Min.X, q.boundary.Max.Y, q.boundary.Min.X, q.boundary.Min.Y, colors[depth%len(colors)])
}

func (q *QuadTree) ForEach(f func(*Particle)) {
	for _, p := range q.particles {
		f(p)
	}

	if q.divided {
		for _, node := range q.nodes {
			node.ForEach(f)
		}
	}
}

func (q *QuadTree) ForBoundary(b Bounds, f func(*Particle)) {
	if !q.boundary.Intersects(b) {
		return
	}

	for _, p := range q.particles {
		if b.Contains(p.Position) {
			f(p)
		}
	}

	if q.divided {
		for _, node := range q.nodes {
			node.ForBoundary(b, f)
		}
	}
}

func (q *QuadTree) subdivide() {
	x := q.boundary.Min.X
	y := q.boundary.Min.Y
	w := q.boundary.Max.X - x
	h := q.boundary.Max.Y - y

	q.divided = true

	q.nodes[0] = NewQuadTree(x, y, x+w/2, y+h/2)
	q.nodes[1] = NewQuadTree(x+w/2, y, x+w, y+h/2)
	q.nodes[2] = NewQuadTree(x, y+h/2, x+w/2, y+h)
	q.nodes[3] = NewQuadTree(x+w/2, y+h/2, x+w, y+h)
}
