package strategy

type Operator interface {
	Apply(int, int) int
}
type Operation struct {
	Operator Operator
}

func (this *Operation) Operate(l, r int) int {
	return this.Operator.Apply(l, r)
}
