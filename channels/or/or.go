package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := func(s time.Duration) chan interface{} {
		c := make(chan interface{})
		go func() {
			time.Sleep(s * time.Second)
			c <- rand.Uint32()
		}()
		return c
	}

	s := time.Now()
	o := or(
		r(3),
		r(5),
		r(2),
		r(6),
		r(1),
		r(4),
	)

	select {
	case i := <-o:
		fmt.Printf("Value received. %d\n", i)
		fmt.Println(time.Since(s))
	}
}

// or receives a slice of channels and returns a single channel.
func or(ch ...<-chan interface{}) <-chan interface{} {
	switch len(ch) {
	case 0:
		return nil
	case 1:
		return ch[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(ch) {
		case 2:
			select {
			case <-ch[0]:
				orDone <- ch[0]
			case <-ch[1]:
				orDone <- ch[1]
			}

		default:
			select {
			// case orDone <-ch[0]:
			// case above won't wait for the channel to return value. Why?
			case <-ch[0]:
				orDone <- ch[0]
			case <-ch[1]:
				orDone <- ch[1]
			case <-ch[2]:
				orDone <- ch[2]
			case <-or(append(ch[3:], orDone)...):
			}
		}
	}()

	return orDone
}
