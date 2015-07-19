package main

import "github.com/nsf/termbox-go"

type Display struct {
	in    chan int
	state DisplayState
	x     int
	y     int
}

func NewDisplay() *Display {
	return &Display{
		state: READX,
	}
}

func (d *Display) Run() {
	go func() {
		termbox.Init()
		termbox.SetOutputMode(termbox.Output256)

		for {
			// x y <colors> -1
			v := <-d.in
			if v == -1 {
				d.state = READX
				continue
			}

			switch d.state {
			case READX:
				d.x = v
				d.state = READY
			case READY:
				d.y = v
				d.state = READCOLOR
			case READCOLOR:
				termbox.SetCell(d.x, d.y, ' ', termbox.ColorDefault, dcolors[v])
				termbox.Flush()
				d.x++
			}
		}
	}()
}

type DisplayState int

const (
	READX DisplayState = iota
	READY
	READCOLOR
)

var dcolors = map[int]termbox.Attribute{
	0: 0x11, // black
	1: 0xed, // dark grey
	2: 0xfc, // bright grey
	3: 0x10, // white
	4: 0xa1, // red
}
