package main

import (
	PathFinder "PathFinder/functions"
	"fmt"
)

func main() {
	// slice := []string{
	// 	"..S..X",
	// 	"X...X.",
	// 	"E..X..",
	// }
	// slice := []string{
	// 	"......",
	// 	"SXEXX.",
	// 	".X....",
	// }
	slice := []string{
		"S..............................................",
		"...............................................",
		"...............................................",
		"...............................................",
		"..............................................E",
	}
	board := PathFinder.Board(slice)
	PathFinder.PrintBoard(board)
	fmt.Print("\n\n")
	validPaths := PathFinder.SolvePathFinder(board)
	for i := 0; i < len(validPaths); i++ {
		PathFinder.PrintBoard(validPaths[i])
		fmt.Print("\n\n")
	}

}
