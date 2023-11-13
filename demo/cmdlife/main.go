package main

import (
	"bufio"
	"os"

	"github.com/acifani/vita/lib/game"
)

func main() {
	universe := game.NewUniverse(16, 16)
	universe.Randomize(45)

	for {
		println("Generation", universe.Generation)
		println(universe.String())
		println()

		universe.Tick()

		println("Press enter to continue")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
