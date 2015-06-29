package main

import "fmt"

func main() {
	n1 := &Node{}
	n2 := &Node{}
	n3 := &Node{}

	n1.ConnectDown(n2)
	n2.ConnectDown(n3)
	ConnectOutput(n3)

	// connect input
	n1.Up = make(chan int)

	n1.Run()
	n2.Run()
	n3.Run()

	for i := 0; i < 10; i++ {
		n1.Up <- i
	}
}

func ConnectOutput(n *Node) {
	n.Down = make(chan int)

	go func() {
		for v := range n.Down {
			fmt.Printf("%d\n", v)
		}
	}()
}
