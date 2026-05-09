package main

import PathFinder "PathFinder/functions"

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

	board := PathFinder.CreateBoard(slice)
	PathFinder.PrintBoard(board)
	replay := true
	for replay {
		replay = PathFinder.UserReplay(board)

	}

	// functions You need "UserReplay" it returns a boolean
	// set a loop if it returns true break out of the loop
	// if it returns stay in the loop
}
