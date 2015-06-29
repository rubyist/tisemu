package main

import "fmt"

func main() {
	n1 := &T21{}
	n2 := &T21{}
	n3 := &T21{}

	n1.ConnectDown(n2)
	n2.ConnectDown(n3)
	ConnectOutput(n3)

	// connect input
	n1.up = make(chan int)

	n1.Run()
	n2.Run()
	n3.Run()

	for i := 0; i < 10; i++ {
		n1.up <- i
	}
}

func ConnectOutput(n *T21) {
	n.down = make(chan int)

	go func() {
		for v := range n.down {
			fmt.Printf("%d\n", v)
		}
	}()
}
