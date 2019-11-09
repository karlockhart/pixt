package geosim

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
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

	node := n.nodeMesh[y][x]
	node.Height = height

	var eval func(*Node) bool
	var action func(*Node)

	if node.Height > node.Prev.Height {
		eval = func(node *Node) bool {
			return node.Prev != nil && node.Height > node.Prev.Height
		}
		action = moveNodeBack
	} else if node.Height < node.Next.Height {
		eval = func(node *Node) bool {
			return node.Next != nil && node.Height < node.Next.Height
		}
		action = moveNodeForward
	}

	i := 0
	for eval(node) {
		action(node)
		i++
	}
	logrus.Debug("moved: ", i)

	for n.MaxSortedHeight.Prev != nil {
		logrus.Debug("mxsh moving", n.MaxSortedHeight.Height)
		n.MaxSortedHeight = n.MaxSortedHeight.Prev
	}
	for n.MinSortedHeight.Next != nil {
		logrus.Debug("mnsh moving", n.MaxSortedHeight.Height)
		n.MinSortedHeight = n.MinSortedHeight.Next
	}
	logrus.Debugf("%v", time.Since(start))
	return nil
}

func moveNodeBack(node *Node) {
	if node.Prev == nil {
		return
	}
	prev := node.Prev
	prev.Next = node.Next
	node.Next = prev
	node.Prev = prev.Prev
	prev.Prev = node
}

func moveNodeForward(node *Node) {
	if node.Next == nil {
		return
	}

	next := node.Next
	next.Prev = node.Prev
	node.Prev = next
	node.Next = next.Next
	next.Next = node
}

func (n *NodeMesh) Meh() {
	fmt.Println("Meh")
}
