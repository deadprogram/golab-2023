package main

import (
	"image/color"
	"time"

	"github.com/acifani/vita/lib/game"
)

const (
	size = 32
)

var (
	universe *game.Universe

	height     uint32 = 32
	width      uint32 = 32
	population        = 55
	gamebuffer []byte

	black = color.RGBA{0, 0, 0, 255}
	white = color.RGBA{255, 255, 255, 255}
)

func main() {
	universe = game.NewUniverse(height, width)
	universe.Randomize(population)

	fullRefreshes := uint(0)
	previousSecond := int64(0)

	gamebuffer = make([]byte, height*width)
	universe.Read(gamebuffer)

	for {
		start := time.Now()

		drawGrid()
		display.Display()

		universe.Read(gamebuffer)
		universe.Tick()

		second := (start.UnixNano() / int64(time.Second))
		if second != previousSecond {
			previousSecond = second
			newFullRefreshes := getFullRefreshes()
			animationTime := time.Since(start)
			animationFPS := int64(10 * time.Second / animationTime)
			print("#", second, " screen=", newFullRefreshes-fullRefreshes, "fps animation=", animationTime.String(), "/", (animationFPS / 10), ".", animationFPS%10, "fps\r\n")
			fullRefreshes = newFullRefreshes
		}
	}
}

func drawGrid() {
	var rows, cols uint32
	c := black

	for rows = 0; rows < height; rows++ {
		for cols = 0; cols < width; cols++ {
			idx := universe.GetIndex(rows, cols)

			switch {
			case universe.Cell(idx) == gamebuffer[idx]:
				// no change, so skip
				continue
			case universe.Cell(idx) == game.Alive:
				c = white
			default: // game.Dead
				c = black
			}

			display.SetPixel(int16(cols), int16(rows), c)
		}
	}
}
