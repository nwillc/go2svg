package main

import (
	"github.com/ajstarks/svgo"
)

type codeMonkey struct {
	canvas *svg.SVG
}

const (
	blackWhite = "stroke: black; stroke-width: 2; fill: white;"
	blackBlack = "stroke: black; fill: black;"
)

func (codeMonkey *codeMonkey) Ear(rx int, ry int) {
	codeMonkey.canvas.Circle(rx, ry, 40, blackWhite)
	codeMonkey.canvas.Circle(rx, ry, 28, blackWhite)
}

func (codeMonkey *codeMonkey) Eye(rx int, ry int) {
	codeMonkey.canvas.Circle(rx, ry, 20, blackWhite)
}

func (codeMonkey *codeMonkey) Nostril(rx int, ry int) {
	codeMonkey.canvas.Circle(rx, ry, 4, blackBlack)
}
