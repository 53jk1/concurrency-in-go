package main

import (
	"github.com/53jk1/concurrency-in-go/ex_context"
)

func main() {
	//queuing.SimplePipeline()
	//queuing.PipelineIncludingBuffer()
	//ex_context.ConcurrentPrintGreetAndFarewell()
	//ex_context.ConcurrentPrintGreetAndFarewellWithContext()
	ex_context.ProcessRequest("jane", "abc123")
}
