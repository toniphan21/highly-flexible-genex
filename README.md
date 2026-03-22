## Highly Flexible Generator Example

This is an example of a highly flexible generator from my [code generation toolbox](https://nhatp.com).

This bundles all the toolbox generators into a single configurable entry point. You have to configure everything to make
it work. The config file looks like this:

```pkl
packages {
  ["github.com/you/project"] {
    chainer {
    // exactly like go-chainer-gen options
    }
    composer {
    // exactly like go-composer-gen options
    }
    stringer {
    // exactly like go-stringer-gen options
    }
  }
}
```

### Requirements

This generator requires the `pkl` binary to be available on your `PATH`. It is used at runtime to evaluate `.pkl`
configuration files. You can install it via:

~~~sh
# macOS
brew install pkl
~~~

For other platforms, see [pkl-lang installation docs](https://pkl-lang.org/main/current/pkl-cli/index.html).

### Usage

Clone the repository and run

~~~sh
# runs the generator against the included example project
go run ./main.go -w ./example

# runs the generator in dry-run mode against the included example project
go run ./main.go --dry-run -w ./example
~~~

### Contributing & License

PRs are welcome! Distributed under the Apache License 2.0.

---

If you like the project, feel free to [buy me a coffee](https://buymeacoffee.com/toniphan21). Thank you!
