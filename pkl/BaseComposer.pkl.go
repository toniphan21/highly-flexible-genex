// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

import (
	"nhatp.com/go/composer-gen/pkl/composer"
	"nhatp.com/go/gen-lib/pkl/gen"
)

type BaseComposer interface {
	composer.Package

	GetOutput() *gen.Output
}

var _ BaseComposer = BaseComposerImpl{}

type BaseComposerImpl struct {
	composer.PackageImpl

	Output *gen.Output `pkl:"output"`
}

func (rcv BaseComposerImpl) GetOutput() *gen.Output {
	return rcv.Output
}
