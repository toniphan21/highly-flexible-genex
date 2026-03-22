package main

import (
	"context"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/alexflint/go-arg"
	"github.com/toniphan21/highly-flexible-genex/pkl"
	"golang.org/x/tools/go/packages"
	chainergen "nhatp.com/go/chainer-gen"
	composergen "nhatp.com/go/composer-gen"
	genlib "nhatp.com/go/gen-lib"
	"nhatp.com/go/gen-lib/cli"
	stringergen "nhatp.com/go/stringer-gen"
)

//go:generate find ./pkl -name "*.pkl.go" -type f -delete
//go:generate pkl-gen-go ./pkl/Config.pkl

type Argument struct {
	WorkingDir     string `arg:"-w,--working-dir" help:"Working directory" default:"." placeholder:"WORKING_DIR"`
	ConfigFileName string `arg:"-c,--config" help:"Config file name" default:"config.pkl" placeholder:"FILE_NAME"`
	DryRun         bool   `arg:"-d,--dry-run" help:"Preview changes without writing to disk"`
}

func (a *Argument) ResolveWorkingDir() string {
	if a.WorkingDir == "" {
		wd, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		return wd
	}

	absPath, err := filepath.Abs(a.WorkingDir)
	if err != nil {
		panic(err)
	}
	return absPath
}

func (a *Argument) ConfigFilePath(defaultName string) string {
	var fn string
	if a.ConfigFileName == "" {
		fn = defaultName
	} else {
		fn = a.ConfigFileName
	}
	return filepath.Join(a.ResolveWorkingDir(), fn)
}

const BinaryVersion = "v0.0.0-example"
const BinaryName = "highly-flexible-genex"
const DefaultConfigFileName = "config.pkl"

func main() {
	var args Argument
	arg.MustParse(&args)

	ctx := context.Background()
	workingDir := args.ResolveWorkingDir()
	configFilePath := args.ConfigFilePath(DefaultConfigFileName)

	if args.DryRun {
		slog.Info(cli.ColorBinary(BinaryName) + " " + cli.ColorVersion(BinaryVersion) + " in DRY mode")
	} else {
		slog.Info(cli.ColorBinary(BinaryName) + " " + cli.ColorVersion(BinaryVersion))
	}
	slog.Info(cli.ColorBinary(BinaryName) + " is working on directory: " + cli.ColorInput(workingDir))
	slog.Info(cli.ColorBinary(BinaryName) + " uses configuration file: " + cli.ColorInput(configFilePath))

	configs, err := pkl.LoadFromPath(ctx, configFilePath)
	if err != nil {
		panic(err)
	}

	fileManager := genlib.NewFileManager(workingDir)

	chainer := chainergen.New(fileManager)
	if cf := configs.ToChainerConfig(); len(cf) > 0 {
		if err = doGenerate(workingDir, fileManager, func(pkg *packages.Package) error {
			return chainer.Generate(pkg, cf)
		}); err != nil {
			panic(err)
		}
	}

	composer := composergen.New(fileManager)
	if cf := configs.ToComposerConfig(); len(cf) != 0 {
		if err = doGenerate(workingDir, fileManager, func(pkg *packages.Package) error {
			return composer.Generate(pkg, cf)
		}); err != nil {
			panic(err)
		}
	}

	stringer := stringergen.New(fileManager)
	if cf := configs.ToStringerConfig(); len(cf) != 0 {
		if err = doGenerate(workingDir, fileManager, func(pkg *packages.Package) error {
			return stringer.Generate(pkg, cf)
		}); err != nil {
			panic(err)
		}
	}

	if args.DryRun {
		slog.Info(cli.ColorBinary(BinaryName) + " is printing generated file content")
		for _, out := range fileManager.Files() {
			cli.PrintFileWithFunction(out.RelPath, []byte(out.Content()), func(l string) {
				slog.Info(l)
			})
		}
	} else {
		slog.Info(cli.ColorBinary(BinaryName) + " is saving generated file to disk")
		for _, out := range fileManager.Files() {
			if err := os.WriteFile(out.FullPath, []byte(out.Content()), 0644); err != nil {
				panic(err)
			}
		}
	}

	slog.Info(cli.ColorGreen("done"))
}

func doGenerate(dir string, fm genlib.FileManager, fn func(p *packages.Package) error) error {
	pkgs, err := genlib.LoadPackagesWithGenFiles(dir, fm)
	if err != nil {
		return err
	}

	for _, pkg := range pkgs {
		if err = fn(pkg); err != nil {
			return err
		}
	}

	return nil
}
