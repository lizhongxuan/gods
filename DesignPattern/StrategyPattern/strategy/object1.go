package strategy
type Addition1 struct{}

func (Addition1) Apply(lval, rval int) int {
	return lval + rval
}