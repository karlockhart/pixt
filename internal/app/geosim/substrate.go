package geosim

const (
	Water = iota
	Lava
	Rock
	Dirt
	Sand
)

// Substrate is a representation of the stack of materials for this node
type Substrate interface {
	GetType() int
	Erode([]Substrate) bool
}


type Rock Struct {

}

type Dirt Struct {

}

type Sand Struct {

}