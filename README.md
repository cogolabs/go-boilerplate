# go-boilerplate
Go boilerplate web server with all Cogo specific best practices

## Table of Contents

- [Go BoilerPlate](#go-boilerplate)
  - [travis](#travis)
  - [cmd](#cmd)
  - [pkg](#pkg)
  - [main.go](#main)
  - [logging](#logging)
  - [metrics](#metrics)
  - [handlers](#handlers)
  - [routers](#routers)
  - [testing](#testing)


# travis
The first thing you're going to want to integrate into your go project is Travis. Travis will build and test your code in a stable environment every time you push a commit to github, ensuring your 
changes will work beyond your local configuration.

To set up Travis you're first going to need a `.travis.yml` file, which you can find on the top level of this directory. Copy what you see there to get the deployment up and running.

# cmd
The `cmd` directory should hold entrypoints into your application.

# pkg
the `pkg` directory should hold the majority of your source code, and anything that you will need to access from inside and only inside your repository.

# main
Any file titled `main.go` represents an entrypoint into your application. This will usually mean two things: the file contains a `main()`, or the file contains a crucial API endpoint.

Note that these files may have command line flags that are parsed using `flag.Parse()`. This will allow you to pass values via the command line, an example being `go run main.go -dsn='foo'` running our `main.go` file with the value `foo` for the variable `dsn`.

In the second case, a `main.go` file will provide an crucial function that will likely be called in some `main()` or by another application. If you are writing a library meant to return a complicated struct, consider adding something like a `New()` constructor in a `main.go` file. This is more so a matter of naming convention, and lower level calls in that package are best kept in other named files.

# types
Golang allows you to construct complicated data types as `structs` like in C, as opposed to `classes` like in Python, Java or other OOP languages.

New type structs should be defined in a standalone package file (see `pkg/mypkg/types.go`), all type definitions in one file. Instance methods on these types should have their own files (see `pkg/mypkg/exampletype.go`), one file for each type definition.

# logging
We use Sentry for all of our logging, see `pkg/observe/logging.go`. The go sentry client library is titled `raven-go`, and we wrap our logs using the `github.com/sirupsen/logrus` library. Using these two tools packages, we can redirect all of our log statements to an associated sentry client. Sentry should only be used for logging exceptions, giving us alerts when anything goes wrong in the produciton or dev environments.

Another note about logrus: the line `import log "github.com/sirupsen/logrus"` will redirect any calls to `log` with calls to `logrus`. This is an example of import naming, something that can be extremely useful if you find two seperate libraries have very similar functionality. Some go libraries (ie. `Logrus`) are deliberately designed as a named import to replace other libraries (ie. `Logging`).

# metrics
We use Prometheus for all of our application metrics, see `pkg/observe/metrics.go`. Just like in Sentry, you're going to need to manually set up your Prometheus connection in your code.

# handlers
All handlers are take in an `http.ResponseWriter` and a `*http.Request` as arguments. Handlers will be called by routers, which will provide a ResponseWriter, and will be the ones receiving http Requests. Handlers are meant to serve content and preform tasks (like a healthcheck, in `pkg/web/handlers.go`) when given these inputs.

# routers
We use `github.com/gorilla/mux` for HTTP routing, as opposed to the default routers provided in the http library. This decision to choose this library for Cogo was made for developer productivity reasons. For applications requiring better performance, prefer the default `net/http` package.

Routers, when instantiated, can call the `handleFunc()` method to call a given handler. You should have one router calling many handlers.

# testing
The best practice for naming tests in golang is `filename_test.go`. These tests should be placed in the same directory, like those under `pkg/web/`. Using this naming convention, all test files will appear right above the file they are testing. Placing your test files in a separate directory is not recommended, unless you are a GoPath wizard and believe that you have the knowledge and reasons to do so.

All golang test functions must start with `Test` in their names, in order for them to be visible to the go tester. Each test should accept the first argument `(t* Testing.t)`, which will most likely be the only argument you need. We use [stretchr/testify](https://github.com/stretchr/testify) as our assertion library to ensure correctness.
