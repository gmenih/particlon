package particle

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

type SpatialIndex struct {
	Boundary   Bounds
	Buckets    [][]*Particle
	width      int
	bucketSize int
	all        []*Particle
}

func NewSpatialIndex(minX, minY, maxX, maxY float64, bucketSize int) *SpatialIndex {
	boundary := BB(minX, minY, maxX, maxY)
	nrBuckets := int(boundary.Max.X*boundary.Max.Y) / bucketSize
	return &SpatialIndex{
		Boundary:   boundary,
		Buckets:    make([][]*Particle, nrBuckets),
		width:      int(boundary.Max.X) / bucketSize,
		bucketSize: bucketSize,
	}
}

func (s *SpatialIndex) Insert(particle *Particle) {
	s.insert(particle)
	s.all = append(s.all, particle)
}

func (s *SpatialIndex) insert(particle *Particle) {
	bucket := s.getBucket(particle.Position)
	if bucket < 0 || bucket >= len(s.Buckets) {
		return
	}
	s.Buckets[bucket] = append(s.Buckets[bucket], particle)
}

func (s *SpatialIndex) Rebalance() {
	s.Buckets = make([][]*Particle, len(s.Buckets))
	for _, p := range s.all {
		s.insert(p)
	}
}

func (s *SpatialIndex) ForEach(f func(*Particle)) {
	for _, p := range s.all {
		f(p)
	}
}

func (s *SpatialIndex) ForAround(position Vector, radius float64, f func(*Particle)) {
	intersections := make(map[int]bool)
	for x := position.X - radius; x <= position.X+radius; x += float64(s.bucketSize) {
		for y := position.Y - radius; y <= position.Y+radius; y += float64(s.bucketSize) {
			bucket := s.getBucket(VV(x, y))
			if bucket < 0 || bucket >= len(s.Buckets) {
				continue
			}
			intersections[bucket] = true
		}
	}

	for bucket := range intersections {
		for _, p := range s.Buckets[bucket] {
			if p.Position.Distance(position) <= radius {
				f(p)
			}
		}
	}
}

func (s *SpatialIndex) getBucket(position Vector) int {
	x := int(math.Floor(position.X / float64(s.bucketSize)))
	y := int(math.Floor(position.Y / float64(s.bucketSize)))

	return x + y*s.width
}

func (s *SpatialIndex) Debug(screen *ebiten.Image) {
	for i, bucket := range s.Buckets {
		if len(bucket) == 0 {
			continue
		}

		color := randomColor(i)
		x := i % s.width
		y := i / s.width

		ebitenutil.DrawLine(screen, float64(x*s.bucketSize), float64(y*s.bucketSize), float64(x*s.bucketSize+s.bucketSize), float64(y*s.bucketSize), color)
		ebitenutil.DrawLine(screen, float64(x*s.bucketSize), float64(y*s.bucketSize), float64(x*s.bucketSize), float64(y*s.bucketSize+s.bucketSize), color)
		ebitenutil.DrawLine(screen, float64(x*s.bucketSize+s.bucketSize), float64(y*s.bucketSize), float64(x*s.bucketSize+s.bucketSize), float64(y*s.bucketSize+s.bucketSize), color)
		ebitenutil.DrawLine(screen, float64(x*s.bucketSize), float64(y*s.bucketSize+s.bucketSize), float64(x*s.bucketSize+s.bucketSize), float64(y*s.bucketSize+s.bucketSize), color)

		for _, p := range bucket {
			ebitenutil.DrawRect(screen, p.Position.X-1, p.Position.Y-1, 2, 2, color)
		}

	}
}

func randomColor(i int) color.Color {
	colors := []color.Color{
		colornames.Red,
		colornames.Orange,
		colornames.Yellow,
		colornames.Blue,
		colornames.Violet,
	}

	return colors[i%len(colors)]
}
