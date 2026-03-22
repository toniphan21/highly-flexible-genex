// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

import "nhatp.com/go/gen-lib/pkl/gen"

type Base interface {
	GetOutput() *gen.Output
}

var _ Base = BaseImpl{}

type BaseImpl struct {
	Output *gen.Output `pkl:"output"`
}

func (rcv BaseImpl) GetOutput() *gen.Output {
	return rcv.Output
}
