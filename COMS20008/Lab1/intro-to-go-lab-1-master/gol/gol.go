package main

import "fmt"

func calculateNextState(p golParams, world [][]byte) [][]byte {
	newWorld := [][]byte{} // new world to store alive cells

	for y := range world {
		row := []byte{}
		for x := range world {
			row = append(row, checkNeighbours(world, y, x))
		}
		newWorld = append(newWorld, row)
	}
	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {
	alive := []cell{}
	for y := range world {
		for x := range world {
			if world[x][y] == 255 {
				alive = append(alive, cell{y, x})
				fmt.Println("calcAlive: ", y, x)
			}
		}
	}
	return alive
}

func checkNeighbours(world [][]byte, x int, y int) byte {
	neighbours := getNeighbours(x, y) // retrieve list of 8 surrounding cells including wrapping
	liveCount := 0                    // number of alive neighbours for given cell

	// updates live count by checking the neighbours values
	for k := range neighbours {
		if world[neighbours[k].x][neighbours[k].y] == 255 {
			liveCount++
		}
	}

	if world[x][y] == 255 {
		if liveCount < 2 || liveCount > 3 { // less than 2 or more than 3 live neighbours cell dies
			return 0
		} else { // else there is 2 or 3 live neighbours so cell lives
			return 255
		}
	} else if liveCount == 3 { // if cell is dead and there are 3 live neighbours then cell lives
		return 255
	} else { // cell is dead and does not have 3 live neighbours so remains dead
		return 0
	}
}

func getNeighbours(x int, y int) []cell {
	xOffset := []int{-1, 0, 1, 1, 1, 0, -1, -1}
	yOffset := []int{1, 1, 1, 0, -1, -1, -1, 0}
	returnVal := []cell{}

	// for each of the 8 neighbours calculate cell locations and offsets
	for v := 0; v < 8; v++ {
		newX := x + xOffset[v]
		newY := y + yOffset[v]
		// to ensure closed grid, aka wrapping around edges
		if newX > 15 {
			newX = 0
		}
		if newX < 0 {
			newX = 15
		}
		if newY > 15 {
			newY = 0
		}
		if newY < 0 {
			newY = 15
		}

		returnVal = append(returnVal, cell{newX, newY})
	}

	return returnVal
}
