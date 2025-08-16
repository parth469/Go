package main

import (
	"fmt"
	"sync"
	"time"
)

func and(channels ...<-chan interface{}) <-chan interface{} {
	orDone := make(chan interface{})

	if len(channels) == 0 {
		close(orDone)
		return orDone
	}

	go func() {
		defer close(orDone)
		for _, ch := range channels {
			<-ch
		}
	}()

	return orDone
}
func or(channels ...<-chan interface{}) <-chan interface{} {
	var once sync.Once
	orDone := make(chan interface{})

	if len(channels) == 0 {
		close(orDone)
		return orDone
	}

	go func() {
		for _, ch := range channels {
			go func(channal <-chan interface{}) {

				select {
				case <-channal:
					once.Do(func() { close(orDone) })
				case <-orDone:
				}
			}(ch)
		}
	}()

	return orDone
}

func Runn(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}
func main() {
	start := time.Now()
	<-or(Runn(2*time.Second),
		Runn(1*time.Second))
	fmt.Printf("done after %v\n", time.Since(start))
	start = time.Now()

	<-and(Runn(6*time.Second),
		Runn(1*time.Second))
	fmt.Printf("done after %v\n", time.Since(start))

	// var or func(channels ...<-chan interface{}) <-chan interface{}
	// or = func(channels ...<-chan interface{}) <-chan interface{} {
	// 	switch len(channels) {
	// 	case 0:
	// 		return nil
	// 	case 1:
	// 		return channels[0]
	// 	}
	// 	orDone := make(chan interface{})
	// 	go func() {
	// 		defer close(orDone)
	// 		switch len(channels) {
	// 		case 2:
	// 			select {
	// 			case <-channels[0]:
	// 			case <-channels[1]:
	// 			}
	// 		default:
	// 			select {
	// 			case <-channels[0]:
	// 			case <-channels[1]:
	// 			case <-channels[2]:
	// 			case <-or(append(channels[3:], orDone)...):
	// 			}
	// 		}
	// 	}()
	// 	return orDone
	// }

}
