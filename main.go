package main

import (
	PathFinder "PathFinder/functions"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
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
	board := [][]string{}
	for {
		for {
			clearScreen()
			userSlice, exit := PathFinder.UserInput()
			if !exit {
				return
			}
			board = PathFinder.CreateBoard(userSlice)
			if !PathFinder.CheckEntireBoard(board) {
				fmt.Println("Your board must contain an E and an S!")
				time.Sleep(2 * time.Second)
			} else {
				break
			}
		}
		// row, col := PathFinder.FindS(board)
		for {
			clearScreen()
			choice, noPath := PathFinder.UserChoice(board)
			if !noPath {
				fmt.Println("Ooops! Sorry There's no light at the end of your tunnel😞😞😞")
				fmt.Println("Just go for deliverance abeg. I would refer you to MFM!")
				time.Sleep(5 * time.Second)
				break
			}
			if choice {
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
