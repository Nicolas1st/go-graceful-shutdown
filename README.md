# Minimalist Graceful Shutdown Implementation in Go

[Example](examples/example.go)

```go
package main

import (
	"fmt"
	shutdown "go-graceful-shutdown"
)

type Api struct {
	s *shutdown.GracefulShutter
}

type Request interface{}

func (api *Api) HandleRequest(r Request) {
	// always must check for the result of the RegOp method
	// if it returns an error, it means it can not register anything anymore
	// the method UnregOp should not be called in this case
	if err := api.s.RegOp(); err == shutdown.ErrFinishedRegistration {
		fmt.Printf("Aborting request %v\n", r)
	} else {
		defer api.s.UnregOp()
		handle(r)
	}
}

func handle(r Request) {
	// all the nesseary ops must be complete
	// by the time this function returns
	fmt.Printf("Finshed handling request %v\n", r)
}

func main() {
	api := Api{s: shutdown.NewGracefulShutter()}

	api.HandleRequest("request1") // prints "Finshed handling request1"

	api.s.Shutdown()

	api.HandleRequest("request2") // prints "Aborting request2"
}

```