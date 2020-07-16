package main

import (
	"github.com/ajstarks/svgo"
)

type codeMonkey struct {
	canvas *svg.SVG
}

const (
	blackBlack = "stroke: black; fill: black;"
	blackWhite = "stroke: black; stroke-width: 2; fill: white;"
)

func (codeMonkey *codeMonkey) ear(rx int, ry int) {
	codeMonkey.canvas.Circle(rx, ry, 40, blackWhite)
	codeMonkey.canvas.Circle(rx, ry, 28, blackWhite)
}

func (codeMonkey *codeMonkey) eye(rx int, ry int) {
	codeMonkey.canvas.Circle(rx, ry, 20, blackWhite)
}

func (codeMonkey *codeMonkey) nostril(rx int, ry int) {
	codeMonkey.canvas.Circle(rx, ry, 4, blackBlack)
}
