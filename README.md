## Table of Contents

- [Overview](#overview)
- [cmd](#cmd)
- [pkg](#pkg)
- [internal](#internal)
- [metrics and pprof](#metrics-and-pprof)
- [web](#web)
- [testing](#testing)
- [Github Actions](#github-actions)

# Overview

Go boilerplate is a simple web server starter pack with Cogo's best practices. These are standard across all projects to ensure easy interoperability, understandability, and discoverability.

To get started, we recommend reading [Go By Example](https://gobyexample.com/), then, after getting more comfortable with the language, [Effective Go](https://go.dev/doc/effective_go). You should also be familiar with [Go Modules](https://go.dev/blog/using-go-modules).

# cmd

The `cmd` directory holds entrypoints into your application. Each binary that gets built and run gets their own directory and `main.go` inside the `cmd` directory.

# pkg

The `pkg` directory holds the majority of your source code. Anything that you will need to access from inside and your repository or other packages lives here.

# internal

The `internal` directory holds your source code that must never be imported by other go modules. The go compiler does not allow imports from `internal` packages that are outside your module.

# metrics and pprof

We use Prometheus for our application metrics, see `pkg/metrics/metrics.go`. You'll need to set up a prometheus scraper to ingest these metrics.

[pprof](https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/) lets you collect useful data on the state of a running go program.

# web

For its rich functionality, we use `github.com/gorilla/mux` for handling HTTP requests.

# testing

Tests are written in files named `*_test.go` and can be run with the `go test` tool.

All golang test functions must start with `Test` in their names, in order for them to be visible to the go tester. Each test should accept the first argument `(t* Testing.t)`, which will most likely be the only argument you need. We use [stretchr/testify](https://github.com/stretchr/testify) as our assertion library to ensure correctness.

# Github Actions

We use Github Actions to compile and test our code on every push. Depending on the project, we may also use it for continuous delivery. The yaml in `.github/workflows/` defines how the code should be built and how the tests should be run.
