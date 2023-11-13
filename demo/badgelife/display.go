package main

import (
	"machine"

	"image/color"
	"time"

	"tinygo.org/x/drivers/st7789"

	"github.com/acifani/vita/lib/game"
)

var (
	display st7789.Device

	gamebuffer []byte

	bk             = color.RGBA{0, 0, 0, 255}
	wh             = color.RGBA{255, 255, 255, 255}
	cellSize int16 = 6
	cellBuf        = []color.RGBA{
		wh, wh, wh, wh, wh, wh,
		wh, wh, bk, bk, wh, wh,
		wh, bk, bk, bk, bk, wh,
		wh, bk, bk, bk, bk, wh,
		wh, wh, bk, bk, wh, wh,
		wh, wh, wh, wh, wh, wh,
	}
)

func startGame() {
	machine.SPI0.Configure(machine.SPIConfig{
		Frequency: 8000000,
		Mode:      0,
	})

	display = st7789.New(machine.SPI0,
		machine.TFT_RST,       // TFT_RESET
		machine.TFT_WRX,       // TFT_DC
		machine.TFT_CS,        // TFT_CS
		machine.TFT_BACKLIGHT) // TFT_LITE

	display.Configure(st7789.Config{
		Rotation: st7789.ROTATION_270,
		Height:   320,
	})

	display.FillScreen(wh)

	gamebuffer = make([]byte, height*width)
	universe.Read(gamebuffer)

	for {
		drawGrid()
		display.Display()
		universe.Read(gamebuffer)

		universe.Tick()

		time.Sleep(10 * time.Millisecond)
	}
}

func drawGrid() {
	var rows, cols uint32

	for rows = 0; rows < height; rows++ {
		for cols = 0; cols < width; cols++ {
			idx := universe.GetIndex(rows, cols)

			switch {
			case universe.Cell(idx) == gamebuffer[idx]:
				// no change, so skip
				continue
			case universe.Cell(idx) == game.Alive:
				display.FillRectangleWithBuffer(1+cellSize*int16(cols), cellSize*int16(rows), cellSize, cellSize, cellBuf)
			default: // game.Dead
				display.FillRectangle(1+cellSize*int16(cols), cellSize*int16(rows), cellSize, cellSize, wh)
			}

		}
	}
}
