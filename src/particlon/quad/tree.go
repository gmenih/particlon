package quad

const (
	NODE_CAPACITY = 16
)

type Identifiable interface {
	Identity() (float64, float64)
}

type Tree[TElement Identifiable] struct {
	boundary Bounds
	elements []TElement
	nodes    [4]*Tree[TElement]

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
		if n.Insert(e) {
			return true
		}
	}

	return false
}

func (t *Tree[TElement]) ForEach(f func(TElement)) {
	if t.isLeaf {
		d := 0
		for i := 0; i < len(t.elements); i++ {
			f(t.elements[i])

			if len(t.elements) > 0 && t.rebalanceElement(t.elements[i], 0) {
				t.elements = append(t.elements[:i], t.elements[i+1-d:]...)
				i--
			}
		}
		return
	}

	for _, n := range t.nodes {
		n.ForEach(f)
	}
}

func (t *Tree[TElement]) rebalanceElement(e TElement, depth int) bool {
	x, y := e.Identity()

	if t.boundary.Contains(x, y) {
		if depth == 0 {
			return false
		}

		return t.Insert(e)
	}

	if t.parent != nil {
		return t.parent.rebalanceElement(e, depth+1)
	}

	return false
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
			n.Insert(e)
		}
	}

	t.elements = nil
}
