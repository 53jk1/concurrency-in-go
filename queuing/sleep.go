package queuing

import "time"

func sleep(done <-chan interface{}, duration time.Duration, values <-chan interface{}) <-chan interface{} {
	sleepStream := make(chan interface{})
	go func() {
		defer close(sleepStream)
		for {
			select {
			case <-done:
				return
			case v := <-values:
				time.Sleep(duration)
				select {
				case <-done:
					return
				case sleepStream <- v:
				}
			}
		}
	}()
	return sleepStream
}
