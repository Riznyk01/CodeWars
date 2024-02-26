package main

import (
	"fmt"
	"sync"
)

func main() {

	a := make(chan string, 3)
	a <- "a-foo"
	a <- "a-bar"
	a <- "a-baz"
	close(a)

	b := make(chan string, 2)
	b <- "b-foo"
	b <- "b-bar"
	close(b)

	c := make(chan string, 2)
	c <- "c-foo"
	c <- "c-bar"
	close(c)

	d := make(chan string, 3)
	d <- "d-foo"
	d <- "d-bar"
	d <- "d-baz"
	close(d)

	merged := MergeMultipleChannels(a, b, c, d)

	for val := range merged {
		fmt.Println(val)
	}
}

func MergeMultipleChannels(c ...chan string) <-chan string {
	out := make(chan string)
	var wg sync.WaitGroup
	wg.Add(len(c))

	processChannel := func(ch <-chan string) {
		defer wg.Done()
		for val := range ch {
			out <- val
		}
	}
	for _, channel := range c {
		go processChannel(channel)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
