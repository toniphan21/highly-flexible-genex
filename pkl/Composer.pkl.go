// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

type Composer interface {
	BaseComposer

	GetPriorities() map[int]BaseComposer
}

var _ Composer = ComposerImpl{}

type ComposerImpl struct {
	BaseComposerImpl

	Priorities map[int]BaseComposer `pkl:"priorities"`
}

func (rcv ComposerImpl) GetPriorities() map[int]BaseComposer {
	return rcv.Priorities
}
