// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

import (
	"nhatp.com/go/chainer-gen/pkl/chainer"
	"nhatp.com/go/gen-lib/pkl/gen"
)

type BaseChainer interface {
	chainer.Package

	GetOutput() *gen.Output
}

var _ BaseChainer = BaseChainerImpl{}

type BaseChainerImpl struct {
	chainer.PackageImpl

	Output *gen.Output `pkl:"output"`
}

func (rcv BaseChainerImpl) GetOutput() *gen.Output {
	return rcv.Output
}
