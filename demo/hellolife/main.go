package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/acifani/vita/lib/game"
)

func main() {
	universe := game.NewUniverse(16, 16)
	universe.Randomize(45)
	println()
	
	for {
		fmt.Println("Generation", universe.Generation)
		fmt.Println(universe)
		fmt.Println()

		universe.Tick()

		fmt.Println("Press enter to continue...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
	}
}
