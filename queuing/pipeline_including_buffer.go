package queuing

import (
	"fmt"
	"time"
)

func PipelineIncludingBuffer() {
	// Create a done channel to signal when the pipeline should stop.
	done := make(chan interface{})
	defer close(done)

	// Create the pipeline stages.
	zeros := take(done, 3, repeat(done, 0))
	shortDelay := delay(done, time.Second, zeros)
	buffer := buffer(done, 2, shortDelay)
	longDelay := delay(done, 4*time.Second, buffer)

	// Process the data received from the last stage of the pipeline.
	for data := range longDelay {
		// Replace this with a more meaningful processing of the data.
		fmt.Println(data)
	}
}

// The `delay` stage sleeps for a given duration before forwarding the data to the next stage.
func delay(done <-chan interface{}, d time.Duration, input <-chan interface{}) <-chan interface{} {
	output := make(chan interface{})

	go func() {
		defer close(output)

		for {
			select {
			case <-done:
				return
			case data, ok := <-input:
				if !ok {
					return
				}

				// Use a timer channel to delay the processing.
				timer := time.NewTimer(d)
				defer timer.Stop()

				select {
				case <-done:
					return
				case <-timer.C:
					// Forward the data to the next stage.
					output <- data
				}
			}
		}
	}()

	return output
}

// The `buffer` stage buffers the data and sends it to the next stage when there are enough elements in the buffer.
func buffer(done <-chan interface{}, size int, input <-chan interface{}) <-chan interface{} {
	output := make(chan interface{})

	go func() {
		defer close(output)

		buffer := make([]interface{}, 0, size)

		for {
			select {
			case <-done:
				return
			case data, ok := <-input:
				if !ok {
					return
				}

				// Add the data to the buffer.
				buffer = append(buffer, data)

				// Send the buffered data to the next stage when there are enough elements in the buffer.
				if len(buffer) == size {
					for _, data := range buffer {
						select {
						case <-done:
							return
						case output <- data:
						}
					}

					// Reset the buffer.
					buffer = buffer[:0]
				}
			}
		}
	}()

	return output
}
