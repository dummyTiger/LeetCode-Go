package test

import "testing"

func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func TestGenerator(t *testing.T) {
	ch := make(chan int)
	go generate(ch)

	for i := 0; i < 20; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}

}
