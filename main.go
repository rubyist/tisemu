package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var inputArgs []string
	var outputArgs []string
	var tisfile string
	maptype := "standard"

	for i := 1; i < len(os.Args); i++ {
		arg := os.Args[i]

		if strings.HasSuffix(arg, ".tis") {
			tisfile = arg
			continue
		}

		if strings.HasPrefix(arg, "-in=") {
			inputArgs = append(inputArgs, arg)
		}

		if strings.HasPrefix(arg, "-out=") {
			outputArgs = append(outputArgs, arg)
		}

		if strings.HasPrefix(arg, "-map=") {
			parts := strings.Split(arg, "=")
			if len(parts) != 2 {
				log.Fatalf("Invalid map type: %s", arg)
			}

			switch parts[1] {
			case "standard", "memory":
				maptype = parts[1]
			default:
				log.Fatalf("Invalid map type: %s", parts[1])
			}
		}
	}

	var tis *Tis100
	switch maptype {
	case "standard":
		tis = NewTis100(StandardMap)
	case "memory":
		tis = NewTis100(MemoryMap)
	}

	for _, arg := range inputArgs {
		in := newInput(arg)
		tis.Input(in.C, in.Node)
		in.Run()
	}

	for _, arg := range outputArgs {
		out := newOutput(arg)
		out.C = tis.Output(out.Node)
		out.Run()
	}

	f, err := os.Open(tisfile)
	if err != nil {
		log.Fatalf("Could not open tisfile: %s", err)
	}
	tis.Program(f)

	tis.Run()

	for {
		time.Sleep(time.Millisecond * 20)
	}
}

type input struct {
	C    chan int
	Node int
	File io.ReadCloser
	name string
}

func newInput(arg string) input {
	// -in=<node>,<file>
	parts := strings.Split(arg, "=")
	if len(parts) != 2 {
		log.Fatalf("Invalid format for arg `%s`", arg)
	}

	nf := strings.Split(parts[1], ",")
	if len(parts) != 2 {
		log.Fatalf("Invalid format for arg `%s`", arg)
	}

	node, err := strconv.Atoi(nf[0])
	if err != nil {
		log.Fatalf("Invalid format for arg `%s`", arg)
	}

	f, err := os.Open(nf[1])
	if err != nil {
		log.Fatalf("Could not open file `%s`: %s", nf[1], err)
	}

	return input{
		C:    make(chan int),
		Node: node,
		File: f,
		name: nf[1],
	}
}

func (i input) Run() {
	go func() {
		scanner := bufio.NewScanner(i.File)
		for scanner.Scan() {
			l := scanner.Text()
			n, err := strconv.Atoi(l)
			if err != nil {
				log.Fatalf("Error parsing line `%s`: %s", l, err)
			}
			i.C <- n
		}
		i.File.Close()
	}()
}

type output struct {
	C    chan int
	Node int
	File *os.File
}

func newOutput(arg string) output {
	// -out=<node>,<file>
	parts := strings.Split(arg, "=")
	if len(parts) != 2 {
		log.Fatalf("Invalid format for arg `%s`", arg)
	}

	nf := strings.Split(parts[1], ",")
	if len(parts) != 2 {
		log.Fatalf("Invalid format for arg `%s`", arg)
	}

	node, err := strconv.Atoi(nf[0])
	if err != nil {
		log.Fatalf("Invalid format for arg `%s`", arg)
	}

	f, err := os.Create(nf[1])
	if err != nil {
		log.Fatalf("Could not open file `%s`: %s", nf[1], err)
	}

	return output{
		Node: node,
		File: f,
	}
}

func (o output) Run() {
	go func() {
		for n := range o.C {
			fmt.Fprintf(o.File, "%d\n", n)
			o.File.Sync()
		}
	}()
}
