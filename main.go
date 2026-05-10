package main

import (
	PathFinder "PathFinder/functions"
	"os"
	"os/exec"
	"runtime"

	"fmt"
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

	board := [][]string{}

	clearScreen()

	// Ask once
	UserName := PathFinder.GetUserName()

	for {
		
		for {

		clearScreen()

		for {

			fmt.Println("================ 🔎 The Pathfinder 🔍 ================")

			userSlice, exit := PathFinder.UserInput(UserName)

			if !exit {
				return
			}

			board = PathFinder.CreateBoard(userSlice)

			if !PathFinder.CheckEntireBoard(board) {

				fmt.Printf("%s, Your board must contain an E and an S now!🙄\n", UserName)
				time.Sleep(1 * time.Second)

			} else {
				break
			}
		}

		for {

			clearScreen()

			fmt.Println("================ 🔎 The Pathfinder 🔍 ================")

			choice, noPath := PathFinder.UserChoice(board)

			if !noPath {

				fmt.Printf("Ooops!🫣 Sorry %s, there's no light at the end of your tunnel😞😞😞\n", UserName)
				fmt.Println("Just go for deliverance abeg.😂 I would refer you to MFM!🤣")

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

			message := fmt.Sprintf("Do you want to replay, %s?🤔 (y/n): \n", UserName)

			replay = strings.ToLower(PathFinder.Input(message))

			if replay == "y" {
				break

			} else if replay == "n" {
				return

			} else {
				fmt.Println("Invalid Input!🤦 Type 'y' for yes or 'n' for no. Simple as ABC! 😂")
			}
		}
	}
}
}
