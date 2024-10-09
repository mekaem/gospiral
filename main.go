package main

import (
	"errors"
	"fmt"
)

func generateSpinnyboi(n int) ([][]rune, error) {
	if n < 0 {
		return nil, errors.New("cannot generate a spinnyboi with negative size")
	}

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
	x, y := 0, 0

	for edge > 0 {
		for j := 0; j < edge; j++ {
			spinnyboi[x][y] = rune('#')
			if j < edge-1 {
				x += directions[direction][0]
				y += directions[direction][1]
			}
		}
		direction = (direction + 1) % 4
		edge--
	}

	return spinnyboi, nil
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
		fmt.Printf("Attempting to print spinnyboi for n of %d\n", n)
		spinnyboi, err := generateSpinnyboi(n)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
			continue
		}
		printSpinnyboi(spinnyboi)
		fmt.Println()
	}
}
