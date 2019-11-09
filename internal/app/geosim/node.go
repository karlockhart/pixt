package geosim

import (
	"fmt"
	"time"
)

const (
	MinHeight = -128
	MaxHeight = 128
)

const (
	NW = iota
	N
	NE
	W
	E
	SW
	S
	SE
)

// Node is a node
type Node struct {
	Neighbor []*Node
	Next     *Node
	Prev     *Node
	Height   int
}

// NodeMesh is
type NodeMesh struct {
	MaxSortedHeight *Node
	MinSortedHeight *Node
	nodeMesh        [][]*Node
}

func newNode() *Node {
	n := Node{}
	n.Height = MinHeight
	n.Neighbor = make([]*Node, 8)
	return &n
}

// NewNodeMesh is
func NewNodeMesh(x, y int) *NodeMesh {
	n := NodeMesh{}
	n.nodeMesh = make([][]*Node, y)

	var prev *Node
	for dy := range n.nodeMesh {
		n.nodeMesh[dy] = make([]*Node, x)
		for dx := range n.nodeMesh[dy] {
			//fmt.Println(dx, dy)
			n.nodeMesh[dy][dx] = newNode()
			if prev != nil {
				n.nodeMesh[dy][dx].Prev = prev
				prev.Next = n.nodeMesh[dy][dx]
			}
			prev = n.nodeMesh[dy][dx]
		}
	}

	for dy := range n.nodeMesh {
		for dx := range n.nodeMesh[dy] {
			if dx-1 >= 0 {
				n.nodeMesh[dy][dx].Neighbor[W] = n.nodeMesh[dy][dx-1]
				if dy-1 >= 0 {
					n.nodeMesh[dy][dx].Neighbor[NW] = n.nodeMesh[dy-1][dx-1]
				}
				if dy+1 < y {
					n.nodeMesh[dy][dx].Neighbor[SW] = n.nodeMesh[dy+1][dx-1]
				}
			}
			if dx+1 < x {
				n.nodeMesh[dy][dx].Neighbor[E] = n.nodeMesh[dy][dx+1]
				if dy-1 >= 0 {
					n.nodeMesh[dy][dx].Neighbor[NE] = n.nodeMesh[dy-1][dx+1]
				}
				if dy+1 < y {
					n.nodeMesh[dy][dx].Neighbor[SE] = n.nodeMesh[dy+1][dx+1]
				}
			}

			if dy-1 >= 0 {
				n.nodeMesh[dy][dx].Neighbor[N] = n.nodeMesh[dy-1][dx]
			}

			if dy+1 < y {
				n.nodeMesh[dy][dx].Neighbor[S] = n.nodeMesh[dy+1][dx]
			}
		}

		n.MaxSortedHeight = n.nodeMesh[0][0]

		if y-1 >= 0 && x-1 >= 0 {
			n.MinSortedHeight = n.nodeMesh[y-1][x-1]
		}
	}

	return &n
}

//SetHeight is
func (n *NodeMesh) SetHeight(x, y, height int) error {
	start := time.Now()
	if y-1 > len(n.nodeMesh) || y-1 < 0 || x-1 > len(n.nodeMesh[y]) || x-1 < 0 || height < -127 || height > 127 {
		return fmt.Errorf("invalid value")
	}
	n.nodeMesh[y][x].Height = height

	i := 0
	for n.nodeMesh[y][x].Prev != nil && n.nodeMesh[y][x].Height > n.nodeMesh[y][x].Prev.Height {
		node := n.nodeMesh[y][x]
		prev := node.Prev
		prev.Next = node.Next
		node.Next = prev
		node.Prev = prev.Prev
		prev.Prev = node
		i++
	}
	fmt.Println("moved: ", i)

	for n.MaxSortedHeight.Prev != nil {
		fmt.Println("mxsh moving", n.MaxSortedHeight.Height)
		n.MaxSortedHeight = n.MaxSortedHeight.Prev
	}
	for n.MinSortedHeight.Next != nil {
		fmt.Println("mnsh moving", n.MaxSortedHeight.Height)
		n.MinSortedHeight = n.MinSortedHeight.Next
	}
	fmt.Printf("%v", time.Since(start))
	return nil
}

func (n *NodeMesh) Meh() {
	fmt.Println("Meh")
}
