package queuing

import (
	"fmt"
	"time"
)

func SimplePipeline() {
	done := make(chan interface{})
	defer close(done)

	zeros := take(done, 3, repeat(done, 0))
	short := sleep(done, 1*time.Second, zeros)
	long := sleep(done, 4*time.Second, short)
	pipeline := long
	fmt.Println("Pipeline:")
	for v := range pipeline {
		fmt.Println(v)
	}
}
