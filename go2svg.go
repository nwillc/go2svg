package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func main() {
	width := 300
	height := 300
	canvas := svg.New(os.Stdout)
	svgCanvas := &CodeMonkey{canvas}
	canvas.Start(width, height)
	// Label
	canvas.Text(40, 50, "#CODE", "font-family:monospace;font-size:40px;")
	// Ears
	svgCanvas.Ear( 100, 100)
	svgCanvas.Ear(240, 70)
	// Face
	canvas.Circle(180, 140, 80, "stroke: black; stroke-width: 2; fill: #aa450f;")
	// Eyes
	svgCanvas.Eye(160, 95)
	svgCanvas.Eye(160+45, 95-5)
	// Muzzle
	canvas.Circle(195, 178, 65, BlackWhite)
	// Nostrils
	svgCanvas.Nostril(178, 138)
	svgCanvas.Nostril(178+35, 138-5)
	// Mouth
	canvas.Path("M 150 150 C 100,250 305,260 230,140 C 205,190 165,170 150,150 Z", "stroke: black; stroke-width: 2; fill: red;")
	canvas.End()
}
