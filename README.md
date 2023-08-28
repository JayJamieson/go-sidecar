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

Get with `go get github.com/JayJamieson/go-sidecar@latest`.

Run `gosidecar init` to create a configuration file.

Simplest usage is to call the `Execute` method providing function name and payload. This will invoke the Lambda
function synchronously until execution completed.

```go
package main

import (
  "context"
  "fmt"
  "os"
)

func main() {
  ctx := context.TODO()
  // gosidecar.New comes from gosidecar package folder.
  // package name can be configured using name in gosidecar.yaml
  goSidecar, _ := gosidecar.New(ctx)

  result, err := goSidecar.Execute(gsc.Image, nil)

  if err != nil {
    fmt.Printf("%v", err)
    os.Exit(1)
  }

  // result is raw Lambda response body, usually encoded json.
  // we can just print it as a string directly for demonstration
  fmt.Printf("%s", string(result))
}
```

## Roadmap

- [ ] Deploy handler functions from configuration
- [ ] Invoke handler functions
  - [ ] Synchronous
  - [ ] Asynchronous
  - [ ] HTTP
