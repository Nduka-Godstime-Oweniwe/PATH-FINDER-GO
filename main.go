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
	
	board := [][]string{}
	username := PathFinder.GetUserName()
	for {
		
		for {
			clearScreen()
			fmt.Println("================ 🔎 The Pathfinder 🔍 ================")
			username := PathFinder.GetUserName()
			userSlice := PathFinder.UserInput()
			board = PathFinder.CreateBoard(userSlice)
			if !PathFinder.CheckEntireBoard(board) {
				fmt.Printf("%s,Your board must contain an E and an S now!🙄\n", username)
				time.Sleep(2 * time.Second)
			} else {
				break
			}
		}
		// row, col := PathFinder.FindS(board)
		for {
			clearScreen()
			fmt.Println("================ 🔎 The Pathfinder 🔍 ================")
			choice, noPath := PathFinder.UserChoice(board)
			if !noPath {
				fmt.Printf("Ooops!🫣 Sorry %s, there's no light at the end of your tunnel😞😞😞\n", username)
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
			message := fmt.Sprintf("Do you want to replay, %s?🤔 (y/n): \n", username)
			replay = strings.ToLower(PathFinder.Input(message))
			if strings.ToLower(replay) == "y" {
				break
			} else if strings.ToLower(replay) == "n" {
				return
			} else {
				println("Invalid Input!🤦 Type 'y' for yes or 'n' for no. Simple as ABC! 😂")
			} 

		}
	}

}
