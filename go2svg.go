package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func ear(canvas *svg.SVG, rx int, ry int) {
	canvas.Circle(rx, ry, 40, "stroke: black; stroke-width: 2; fill: white;")
	canvas.Circle(rx, ry, 28, "stroke: black; stroke-width: 2; fill: white;")
}

func eye(canvas *svg.SVG, rx int, ry int) {
	canvas.Circle(rx, ry, 20, "stroke: black; stroke-width: 2; fill: white;")
}

func nostril(canvas *svg.SVG, rx int, ry int) {
	canvas.Circle(rx, ry, 4, "stroke: black; fill: black;")
}

func main() {
	width := 300
	height := 300
	canvas := svg.New(os.Stdout)
	canvas.Start(width, height)
	// Label
	canvas.Text(40, 50, "#CODE", "font-family:monospace;font-size:40px;")
	// Ears
	ear(canvas, 100, 100)
	ear(canvas, 240, 70)
	// Face
	canvas.Circle(180, 140, 80, "stroke: black; stroke-width: 2; fill: #aa450f;")
	// Eyes
	eye(canvas, 160, 95)
	eye(canvas, 160+45, 95-5)
	// Muzzle
	canvas.Circle(195, 178, 65, "stroke: black; stroke-width: 2; fill: white;")
	// Nostrils
	nostril(canvas, 178, 138)
	nostril(canvas, 178+35, 138-5)
	// Mouth
	canvas.Path("M 150 150 C 100,250 305,260 230,140 C 205,190 165,170 150,150 Z", "stroke: black; stroke-width: 2; fill: red;")
	canvas.End()
}
