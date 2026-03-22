// Code generated from Pkl module `genex.Config`. DO NOT EDIT.
package pkl

import (
	"context"

	"github.com/apple/pkl-go/pkl"
	"nhatp.com/go/chainer-gen/pkl/chainer"
)

type Config struct {
	DefaultChainedMethods []chainer.ChainedMethod `pkl:"defaultChainedMethods"`

	All Base `pkl:"all"`

	Packages map[string]Package `pkl:"packages"`
}

// LoadFromPath loads the pkl module at the given path and evaluates it into a Config
func LoadFromPath(ctx context.Context, path string) (ret Config, err error) {
	evaluator, err := pkl.NewEvaluator(ctx, pkl.PreconfiguredOptions)
	if err != nil {
		return ret, err
	}
	defer func() {
		cerr := evaluator.Close()
		if err == nil {
			err = cerr
		}
	}()
	ret, err = Load(ctx, evaluator, pkl.FileSource(path))
	return ret, err
}

// Load loads the pkl module at the given source and evaluates it with the given evaluator into a Config
func Load(ctx context.Context, evaluator pkl.Evaluator, source *pkl.ModuleSource) (Config, error) {
	var ret Config
	err := evaluator.EvaluateModule(ctx, source, &ret)
	return ret, err
}
