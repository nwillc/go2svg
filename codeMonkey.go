package main

import (
	"github.com/ajstarks/svgo"
)

type CodeMonkey struct {
	*svg.SVG
}

const (
	BlackWhite = "stroke: black; stroke-width: 2; fill: white;"
	BlackBlack = "stroke: black; fill: black;"
)

func (canvas *CodeMonkey) Ear(rx int, ry int) {
	canvas.SVG.Circle(rx, ry, 40, BlackWhite)
	canvas.SVG.Circle(rx, ry, 28, BlackWhite)
}

func (canvas *CodeMonkey) Eye(rx int, ry int) {
	canvas.Circle(rx, ry, 20, BlackWhite)
}

func (canvas *CodeMonkey) Nostril(rx int, ry int) {
	canvas.Circle(rx, ry, 4, BlackBlack)
}
