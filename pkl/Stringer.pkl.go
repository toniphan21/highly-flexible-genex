// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

type Stringer interface {
	BaseStringer

	GetPriorities() map[int]BaseStringer
}

var _ Stringer = StringerImpl{}

type StringerImpl struct {
	BaseStringerImpl

	Priorities map[int]BaseStringer `pkl:"priorities"`
}

func (rcv StringerImpl) GetPriorities() map[int]BaseStringer {
	return rcv.Priorities
}
