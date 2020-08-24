package strategy
type Addition2 struct{}

func (Addition2) Apply(lval, rval int) int {
	return lval * rval
}