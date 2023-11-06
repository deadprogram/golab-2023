package main

import (
	"image/color"
	"sync"
	"time"

	"github.com/acifani/vita/lib/game"
)

var (
	multiverse []*game.ParallelUniverse
	multirows  = 2
	multicols  = 3

	height     uint32 = 32
	width      uint32 = 32
	population        = 35
	gamebuffers [][]byte

	dead = color.RGBA{0, 0, 0, 255}
	alive = color.RGBA{0, 255, 0, 255}
)

func main() {
	multiverse = createUniverses()
	connectUniverses(multiverse)

	for i := 0; i < len(multiverse); i++ {
		gamebuffers = append(gamebuffers, make([]byte, height*width))
		multiverse[i].Read(gamebuffers[i])
	}

	fullRefreshes := uint(0)
	previousSecond := int64(0)

	for {
		start := time.Now()

		drawCube()
		display.Display()

		for i := 0; i < len(multiverse); i++ {
			multiverse[i].Read(gamebuffers[i])
		}
	
		runUniverses(multiverse)

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

func drawCube()	{
	for i := range multiverse {
		drawSide(int16(i), multiverse[i], gamebuffers[i])
	}
}

func drawSide(side int16, u *game.ParallelUniverse, gamebuffer []byte) {
	var rows, cols uint32
	c := dead

	for rows = 0; rows < height; rows++ {
		for cols = 0; cols < width; cols++ {
			idx := u.GetIndex(rows, cols)

			switch {
			case u.Cell(idx) == gamebuffer[idx]:
				// no change, so skip
				continue
			case u.Cell(idx) == game.Alive:
				c = alive
			default: // game.Dead
				c = dead
			}

			display.SetPixel(int16(cols)+side*int16(width), int16(rows), c)
		}
	}
}

func createUniverses() []*game.ParallelUniverse {
	multi := []*game.ParallelUniverse{}
	for i := 0; i < 6; i++ {
		u := game.NewParallelUniverse(height, width)
		u.Randomize(population)
		multi = append(multi, u)
	}

	return multi
}

func connectUniverses(multi []*game.ParallelUniverse) {
	multi[0].SetTopNeighbor(multi[3])
	multi[0].SetRightNeighbor(multi[1])
	multi[0].SetBottomNeighbor(multi[3])
	multi[0].SetLeftNeighbor(multi[2])

	multi[1].SetTopNeighbor(multi[4])
	multi[1].SetRightNeighbor(multi[2])
	multi[1].SetBottomNeighbor(multi[4])

	multi[2].SetTopNeighbor(multi[5])
	multi[2].SetBottomNeighbor(multi[5])

	multi[3].SetLeftNeighbor(multi[5])
	multi[3].SetRightNeighbor(multi[4])

	multi[4].SetRightNeighbor(multi[5])
}

func runUniverses(multi []*game.ParallelUniverse) {
	var wg sync.WaitGroup
	for _, u := range multi {
		callMultiTick(&wg, u)
	}
	wg.Wait()
}

func callMultiTick(wg *sync.WaitGroup, u *game.ParallelUniverse) {
	wg.Add(1)
	go func() {
		u.MultiTick()
		wg.Done()
	}()
}
