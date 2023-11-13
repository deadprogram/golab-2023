package main

import (
	"machine"

	"time"

	"github.com/acifani/vita/lib/game"
)

var (
	universe *game.Universe

	height     uint32 = 40
	width      uint32 = 53
	population        = 20
)

func main() {
	universe = game.NewUniverse(height, width)
	universe.Randomize(population)

	go startGame()

	btnA := machine.BUTTON_A
	btnB := machine.BUTTON_B
	btnA.Configure(machine.PinConfig{Mode: machine.PinInput})
	btnB.Configure(machine.PinConfig{Mode: machine.PinInput})

	for {
		if !btnA.Get() {
			universe.Randomize(population)
		}

		if !btnB.Get() {
			universe.Reset()
		}

		time.Sleep(time.Millisecond * 200)
	}
}
