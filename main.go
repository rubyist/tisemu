package main

import (
	"fmt"
	"strings"
)

func main() {
	m := NodeMap{
		{T21Node, T21Node, T21Node, T21Node},
		{T21Node, T21Node, T21Node, T21Node},
		{T21Node, T21Node, T21Node, T21Node},
	}

	tis := NewTis100(m)

	in := make(chan int)
	tis.Input(in, 1)
	out := tis.Output(10)

	tis.Program(strings.NewReader(p))
	tis.Run()

	go func() {
		for i := 0; i < 20; i++ {
			in <- i
		}
	}()

	for i := 0; i < 20; i++ {
		fmt.Println(<-out)
	}
}

var p = `@0

@1
MOV UP ACC
MOV ACC RIGHT
ADD RIGHT
MOV ACC DOWN

@2
MOV LEFT LEFT

@3

@4

@5
MOV UP DOWN

@6

@7

@8

@9
MOV UP RIGHT

@10
MOV LEFT DOWN

@11
`
