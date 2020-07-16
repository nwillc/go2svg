package main

import (
	"github.com/ajstarks/svgo"
	"os"
)

func main() {
	codeMonkey := &CodeMonkey{svg.New(os.Stdout)}
	codeMonkey.canvas.Start(300, 300)
	// Label
	codeMonkey.canvas.Text(40, 50, "#CODE", "font-family:monospace;font-size:40px;")
	// Ears
	codeMonkey.ear(100, 100)
	codeMonkey.ear(240, 70)
	// Face
	codeMonkey.canvas.Circle(180, 140, 80, "stroke: black; stroke-width: 2; fill: #aa450f;")
	// Eyes
	codeMonkey.eye(160, 95)
	codeMonkey.eye(160+45, 95-5)
	// Muzzle
	codeMonkey.canvas.Circle(195, 178, 65, blackWhite)
	// Nostrils
	codeMonkey.nostril(178, 138)
	codeMonkey.nostril(178+35, 138-5)
	// Mouth
	codeMonkey.canvas.Path("M 150 150 C 100,250 305,260 230,140 C 205,190 165,170 150,150 Z", "stroke: black; stroke-width: 2; fill: red;")
	codeMonkey.canvas.End()
}
