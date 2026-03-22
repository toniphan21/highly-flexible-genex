// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

import (
	"nhatp.com/go/gen-lib/pkl/gen"
	"nhatp.com/go/stringer-gen/pkl/stringer"
)

type BaseStringer interface {
	stringer.Package

	GetOutput() *gen.Output
}

var _ BaseStringer = BaseStringerImpl{}

type BaseStringerImpl struct {
	stringer.PackageImpl

	Output *gen.Output `pkl:"output"`
}

func (rcv BaseStringerImpl) GetOutput() *gen.Output {
	return rcv.Output
}
