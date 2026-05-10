package PathFinder

import (
	"os"
	"os/exec"
	"runtime"
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

// Creates The Board
func CreateBoard(slice []string) [][]string {
	empty := [][]string{}
	str := []string{}
	for i := 0; i < len(slice); i++ {
		for _, v := range slice[i] {
			str = append(str, string(v))
		}
		empty = append(empty, str)
		str = []string{}
	}
	return empty
}

func CheckEntireBoard(board [][]string) bool {
	countS := 0
	countE := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "S" {
				countS++
			}

			if board[i][j] == "E" {
				countE++
			}
		}
	}

	if countE != 1 && countS != 1 {
		return false
	}
	return true
}
