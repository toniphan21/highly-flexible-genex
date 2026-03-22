// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

type Chainer interface {
	BaseChainer

	GetPriorities() map[int]BaseChainer
}

var _ Chainer = ChainerImpl{}

type ChainerImpl struct {
	BaseChainerImpl

	Priorities map[int]BaseChainer `pkl:"priorities"`
}

func (rcv ChainerImpl) GetPriorities() map[int]BaseChainer {
	return rcv.Priorities
}
