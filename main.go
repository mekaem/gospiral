package main

import (
	"fmt"
)

func generateSpinnyboi(n int) [][]rune {
	spinnyboi := make([][]rune, n)
	for i := range spinnyboi {
		spinnyboi[i] = make([]rune, n)
		for j := range spinnyboi[i] {
			spinnyboi[i][j] = ' '
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	direction := 0
	edge := n
	x := 0
	y := 0

	for count := 0; edge > 0; edge-- {
		for j := 0; j < edge; j++ {
			spinnyboi[x][y] = rune('#')
			if j < edge-1 {
				x += directions[direction][0]
				y += directions[direction][1]
			}
		}
		count++
		direction = (direction + 1) % 4
	}
	return spinnyboi
}

func printSpinnyboi(spinnyboi [][]rune) {
	for _, row := range spinnyboi {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func main() {
	for n := 1; n <= 7; n++ {
		fmt.Println("Printing spinnyboi for n of", n)
		spinnyboi := generateSpinnyboi(n)
		printSpinnyboi(spinnyboi)
	}
}
