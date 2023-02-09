package quad

import "gmenih/particlon/src/particlon/base"

const (
	NODE_CAPACITY = 16
)

type Identifiable interface {
	Identity() base.Vector
}

type Tree[TElement Identifiable] struct {
	boundary    Bounds
	elements    []TElement
	allElements []TElement
	nodes       [4]*Tree[TElement]

	isLeaf   bool
	capacity int
	parent   *Tree[TElement]
}

func NewTree[TElement Identifiable](boundary Bounds, parent *Tree[TElement]) *Tree[TElement] {
	return &Tree[TElement]{
		elements: make([]TElement, 0, NODE_CAPACITY),
		nodes:    [4]*Tree[TElement]{},
		boundary: boundary,
		isLeaf:   true,
		capacity: NODE_CAPACITY,
		parent:   parent,
	}
}

func (t *Tree[TElement]) Insert(e TElement) bool {
	v := t.insert(e)
	if v && t.parent == nil {
		t.allElements = append(t.allElements, e)
	}

	return v
}

func (t *Tree[TElement]) insert(e TElement) bool {
	if !t.boundary.Contains(e.Identity()) {
		return false
	}

	if t.isLeaf {
		if len(t.elements) < t.capacity {
			t.elements = append(t.elements, e)
			return true
		}

		t.split()
	}

	for _, n := range t.nodes {
		if n.insert(e) {
			return true
		}
	}

	return false
}

func (t *Tree[TElement]) ForEach(f func(TElement)) {
	for _, e := range t.allElements {
		f(e)
	}
}

func (t *Tree[TElement]) QueryRange(b Bounds) []TElement {
	if !t.boundary.Intersects(b) {
		return nil
	}

	var elements []TElement
	if t.isLeaf {
		for _, e := range t.elements {
			if b.Contains(e.Identity()) {
				elements = append(elements, e)
			}
		}

		return elements
	} else {
		for _, n := range t.nodes {
			elements = append(elements, n.QueryRange(b)...)
		}
	}

	return elements
}

func (t *Tree[TElement]) split() {
	t.isLeaf = false

	x := t.boundary.MinX
	y := t.boundary.MinY
	w := t.boundary.MaxX - t.boundary.MinX
	h := t.boundary.MaxY - t.boundary.MinY

	t.nodes = [4]*Tree[TElement]{
		NewTree(NewBounds(x, y, x+w/2, y+h/2), t),
		NewTree(NewBounds(x+w/2, y, x+w, y+h/2), t),
		NewTree(NewBounds(x, y+h/2, x+w/2, y+h), t),
		NewTree(NewBounds(x+w/2, y+h/2, x+w, y+h), t),
	}

	for _, e := range t.elements {
		for _, n := range t.nodes {
			n.insert(e)
		}
	}

	t.elements = nil
}

func (t *Tree[TElement]) Rebalance() {
	t.nodes = [4]*Tree[TElement]{}
	t.isLeaf = true
	t.elements = make([]TElement, 0, NODE_CAPACITY)

	for _, e := range t.allElements {
		t.insert(e)
	}
}
