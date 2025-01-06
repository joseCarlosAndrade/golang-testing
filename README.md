# Using mockery

Installing:

```bash
    go install github.com/vektra/mockery/v2@latest
    mockery --version # assert mockery was installed
```

Have a directory tree like this:

```text
/your-go-project
  ├── go.mod
  ├── service/
  │   ├── service.go
  │   └── service_test.go
```

Create an interface in `service.go` (check out the `mocking/mocking.go` file) with the desired methods.

## Creating mocks

On the project root, run:

```bash
mockery --name=InterfaceName -recursive --output=./mocks
```

This will mock the `InterfaceName` interface, searching for it recursively. It will put the mock inside the `mocks` package (directory).

## Creating the unit test

On `service_test.go` (your test file) import the following:

```go
import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock" // may not be used
	"your-project-name/mocks"
)
```

Create the unit test

```go
func TestMethodFromInterface(t * testing.T) {
    mockMethodFromInterface := new(mocks.InterfaceName)

    // mock the MethodName when called with argument 1
    mockMethodFromInterface.On("MehodName", 1).Return("return statement 1", "return statement 2", nil)

    return1, return2, err := mockMethodFromInterface(1)

    // do whatever you want with the results
    assert.NoError(t, err)

    assert.Equal(t, "return statement 1", return1) // ok
    assert.Equal(t, "return statement 20", return2) // error
    
    mockMethodFromInterface.AssertExpectations(t) // runs all the assertions
}
```

## Running the tests:

```bash
go test ./...

```

## Makefile

Use `make mocks` for mocks creation automation.
