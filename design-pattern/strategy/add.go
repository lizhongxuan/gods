package main
type Add struct{}

func (Add) Apply(lval, rval int) int {
	return lval + rval
}