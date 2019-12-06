package geosim

type Fluid interface {
	Flow([]Substrate, []*Node)
}
