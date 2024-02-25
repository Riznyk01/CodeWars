package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan string, 3)
	a <- "foo"
	a <- "bar"
	a <- "baz"
	close(a)

	b := make(chan string, 2)
	b <- "hello"
	b <- "world"
	close(b)

	c := Merge(a, b)

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

}
func Merge(a <-chan string, b <-chan string) <-chan string {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	processChannel := func(ch <-chan string) {
		defer wg.Done()
		for val := range ch {
			c <- val
		}
	}

	go processChannel(a)
	go processChannel(b)

	go func() {
		wg.Wait()
		close(c)
	}()
	return c
}
