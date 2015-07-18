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
# Signal Amplifier
MOV UP ANY

@2
MOV LEFT ACC
ADD ACC
MOV ACC DOWN

@3

@4

@5
MOV ANY ACC
ADD ACC
MOV ACC RIGHT

@6
MOV ANY DOWN

@7

@8

@9

@10
MOV UP DOWN

@11
`
