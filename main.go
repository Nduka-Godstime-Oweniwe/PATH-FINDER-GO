package main

import (
	"fmt"

	PathFinder "PathFinder/functions"
)

func main() {
	slice := []string{
		"..S..X",
		"X...X.",
		"E..X..",
	}
	// slice := []string{
	// 	"......",
	// 	"SXEXX.",
	// 	".X....",
	// // }
	// slice := []string{
	// 	"S..............................................",
	// 	"...............................................",
	// 	"...............................................",
	// 	"...............................................",
	// 	"..............................................E",
	// }
	board := PathFinder.Board(slice)
	fmt.Println("YOUR GRID:")
	PathFinder.PrintBoard(board)

	validPaths := PathFinder.SolvePathFinder(board)
	if len(validPaths) == 0 {
		fmt.Println("Ooops! There are no valid Path.")
		return

	}
	paths := PathFinder.Choice()
	if paths == "1" {
		count := PathFinder.CountS(validPaths[0])
		paths := 1
		for i := 0; i < len(validPaths); i++ {
			if count == PathFinder.CountS(validPaths[i]) {
				fmt.Printf("Path %d:\n", paths)
				paths++
				PathFinder.PrintBoard(validPaths[i])
				fmt.Print("\n\n")
			}
		}

	} else if paths == "2" {
		count := PathFinder.CountS(validPaths[len(validPaths)-1])
		paths := 1
		for i := 0; i < len(validPaths); i++ {
			if count == PathFinder.CountS(validPaths[i]) {
				fmt.Printf("Path %d:\n", paths)
				paths++
				PathFinder.PrintBoard(validPaths[i])
				fmt.Print("\n\n")
			}
		}
	} else {
		paths := 1
		for i := 0; i < len(validPaths); i++ {
			fmt.Printf("Path %d:\n", paths)
			paths++
			PathFinder.PrintBoard(validPaths[i])
			fmt.Print("\n\n")
		}
	}

}
