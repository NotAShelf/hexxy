package main

import (
	"strconv"
)

const GREY = "\x1b[38;2;111;111;111m"
const CLR = "\x1b[0m"

var ESC = []byte{0x5c, 0x78, 0x31, 0x62, 0x5b}
var CLEAR = []byte{0x5c, 0x78, 0x31, 0x62, 0x5b, 0x30, 0x6d}
var CLRR = []byte("\x1b[0m")

type Color struct {
	disable bool
	values  [256][]byte
}

func (c *Color) Compute() {
	const WHITEB = "\x1b[1;37m"
	for i := 0; i < 256; i++ {
		var fg, bg []byte

		lowVis := i == 0 || (i >= 16 && i <= 20) || (i >= 232 && i <= 242)

		if lowVis {
			fg = append([]byte(WHITEB), []byte("\x1b[38;5;255m")...)
			bg = []byte("\x1b[48;5;" + strconv.Itoa(i) + "m")
		} else {
			fg = []byte("\x1b[38;5;" + strconv.Itoa(i) + "m")
			bg = nil
		}
		c.values[i] = append(bg, fg...)
	}
}

func (c *Color) Colorize(s string, clr byte) string {
	const NOCOLOR = "\x1b[0m"
	colorCode := c.values[clr]
	return string(append(append(colorCode, s...), []byte(NOCOLOR)...))
}

// function to colorize bytes - avoiding string conversions
func (c *Color) Colorize2(clr byte) ([]byte, []byte) {
	colorCode := c.values[clr]
	return colorCode, CLRR
}
