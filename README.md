# Go-Sidecar (WIP)

**Work in progress** not a working release yet.

Somewhat of a fork of [Sidecar](https://github.com/hammerstonedev/sidecar). I wanted something similar but without
the reliance on PHP. 

**CLI Installation**

`go install github.com/JayJamieson/go-sidecar@latest`

## Aim

[Sidecar](https://github.com/hammerstonedev/sidecar) is a good library, but it doesn't come with an
HTTP interface for invoking Lambda functions directly. This may not be a problem if you're working in
PHP land. Go-Sidecar aims to allow better reuse of those sidecar Lambdas from multiple languages.

There are two options for this:

- API Gateway with single invoke endpoint.
  - Probably if you don't use API Gateway for anything else
- Utilize [Lambda function URLs](https://docs.aws.amazon.com/lambda/latest/dg/lambda-urls.html)
  - Less heavy than full-blown API Gateway

## Usage

Install with `go get github.com/JayJamieson/go-sidecar@latest`.

Simplest usage is to call the `Invoke` method providing function name and payload. This will invoke the Lambda
function synchronously until execution completed.

Run `gosidecar init` to create a configuration file.

```go
package main

import "github.com/JayJamieson/go-sidecar"

func main() {
	gosidecar.Invoke("foo", struct {
		foo string
		bar int
	}{
		foo: "bar",
		bar: 42,
    })
}
```

## Roadmap

- [ ] Deploy handler functions from configuration
- [ ] Invoke handler functions
  - [ ] Synchronous
  - [ ] Asynchronous
  - [ ] HTTP