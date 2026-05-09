package main

import (
	PathFinder "PathFinder/functions"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
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
	// // }
	// slice := []string{
	// 	"S..............................................",
	// 	"...............................................",
	// 	"...............................................",
	// 	"...............................................",
	// 	"..............................................E",
	// }
	for {
		clearScreen()
		userSlice := PathFinder.UserInput()
		board := PathFinder.CreateBoard(userSlice)
		// row, col := PathFinder.FindS(board)
		for {
			clearScreen()
			if PathFinder.UserChoice(board) {
				continue
			} else {
				break
			}
		}
		replay := ""
		for replay != "n" {
			replay = strings.ToLower(PathFinder.Input("Do you want to replay? (y/n): "))
			if strings.ToLower(replay) == "y" {
				break
			} else if strings.ToLower(replay) == "n" {
				return
			} else {
				println("Invalid Input. Type 'y' for yes or 'n' for no.")
			}

		}
	}

}
