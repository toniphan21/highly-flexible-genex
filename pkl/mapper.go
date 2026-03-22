package pkl

import (
	"slices"
	"sort"
	"strings"

	chainergen "nhatp.com/go/chainer-gen"
	"nhatp.com/go/chainer-gen/pkl/chainer"
	composergen "nhatp.com/go/composer-gen"
	"nhatp.com/go/composer-gen/pkl/composer"
	"nhatp.com/go/gen-lib"
	"nhatp.com/go/gen-lib/pkl/gen"
	stringergen "nhatp.com/go/stringer-gen"
	"nhatp.com/go/stringer-gen/pkl/stringer"
)

func mergeOutputs(c *Config, pkg Package, component *gen.Output) genlib.Output {
	output := c.All.GetOutput().ToOutput()

	po := pkg.Output
	if po != nil {
		output.Merge(po.ToOutput())
	}

	if component != nil {
		output.Merge(component.ToOutput())
	}

	return output
}

// --- chainer

func makeChainerConfig(c *Config, pkgPath string, pkg Package, base BaseChainer) (*chainergen.Config, error) {
	cf, err := chainer.ToConfig(base)
	if err != nil || cf == nil {
		return nil, err
	}

	cf.Output = mergeOutputs(c, pkg, base.GetOutput())
	cf.PackagePath = pkgPath
	return cf, nil
}

func (c *Config) ToChainerConfig() []chainergen.Config {
	var result []chainergen.Config
	for pkgPath, pkg := range c.Packages {
		priorities := pkg.Chainer.GetPriorities()
		if priorities != nil {
			var priorityKeys []int
			for i := range priorities {
				priorityKeys = append(priorityKeys, i)
			}
			sort.Ints(priorityKeys)

			for _, i := range priorityKeys {
				p := priorities[i]
				priorityCf, err := makeChainerConfig(c, pkgPath, pkg, p)
				if err != nil || priorityCf == nil {
					continue
				}
				result = append(result, *priorityCf)
			}
		}

		defaultCf, err := makeChainerConfig(c, pkgPath, pkg, pkg.Chainer)
		if err != nil || defaultCf == nil {
			continue
		}
		result = append(result, *defaultCf)
	}

	slices.SortFunc(result, func(a, b chainergen.Config) int {
		return strings.Compare(a.PackagePath, b.PackagePath)
	})
	return result
}

// --- composer

func makeComposerConfig(c *Config, pkgPath string, pkg Package, base BaseComposer) (*composergen.Config, error) {
	cf, err := composer.ToConfig(base)
	if err != nil || cf == nil {
		return nil, err
	}
	cf.Output = mergeOutputs(c, pkg, base.GetOutput())
	cf.PackagePath = pkgPath
	return cf, nil
}

func (c *Config) ToComposerConfig() []composergen.Config {
	var result []composergen.Config
	for pkgPath, pkg := range c.Packages {
		priorities := pkg.Composer.GetPriorities()
		if priorities != nil {
			var priorityKeys []int
			for i := range priorities {
				priorityKeys = append(priorityKeys, i)
			}
			sort.Ints(priorityKeys)

			for _, i := range priorityKeys {
				p := priorities[i]
				priorityCf, err := makeComposerConfig(c, pkgPath, pkg, p)
				if err != nil || priorityCf == nil {
					continue
				}
				result = append(result, *priorityCf)
			}
		}

		defaultCf, err := makeComposerConfig(c, pkgPath, pkg, pkg.Composer)
		if err != nil || defaultCf == nil {
			continue
		}
		result = append(result, *defaultCf)
	}

	slices.SortFunc(result, func(a, b composergen.Config) int {
		return strings.Compare(a.PackagePath, b.PackagePath)
	})
	return result
}

// --- stringer

func makeStringerConfig(c *Config, pkgPath string, pkg Package, base BaseStringer) (*stringergen.Config, error) {
	cf, err := stringer.ToConfig(base)
	if err != nil || cf == nil {
		return nil, err
	}
	cf.Output = mergeOutputs(c, pkg, base.GetOutput())
	cf.PackagePath = pkgPath
	return cf, nil
}

func (c *Config) ToStringerConfig() []stringergen.Config {
	var result []stringergen.Config
	for pkgPath, pkg := range c.Packages {
		priorities := pkg.Stringer.GetPriorities()
		if priorities != nil {
			var priorityKeys []int
			for i := range priorities {
				priorityKeys = append(priorityKeys, i)
			}
			sort.Ints(priorityKeys)

			for _, i := range priorityKeys {
				p := priorities[i]
				priorityCf, err := makeStringerConfig(c, pkgPath, pkg, p)
				if err != nil || priorityCf == nil {
					continue
				}
				result = append(result, *priorityCf)
			}
		}

		defaultCf, err := makeStringerConfig(c, pkgPath, pkg, pkg.Stringer)
		if err != nil || defaultCf == nil {
			continue
		}
		result = append(result, *defaultCf)
	}

	slices.SortFunc(result, func(a, b stringergen.Config) int {
		return strings.Compare(a.PackagePath, b.PackagePath)
	})
	return result
}
