package PathFinder

import "fmt"

func Change_S_To_Arrow(board [][]string, rowStart, colStart int) [][]string {
	// CheckUp
	if rowStart != 0 {
		if board[rowStart-1][colStart] == "E" {
			return board
		} else if board[rowStart-1][colStart] == "S" {
			board[rowStart][colStart] = "⬆️"
			Change_S_To_Arrow(board, rowStart-1, colStart)
		}
	}

	//CheckDown
	if rowStart != len(board)-1 {
		if board[rowStart+1][colStart] == "E" {
			return board
		} else if board[rowStart+1][colStart] == "S" {
			board[rowStart][colStart] = "⬇️"
			Change_S_To_Arrow(board, rowStart+1, colStart)
		}
	}

	//CheckRight
	if colStart != len(board[0])-1 {
		if board[rowStart][colStart+1] == "E" {
			return board
		} else if board[rowStart][colStart+1] == "S" {
			board[rowStart][colStart] = "➡️"
			Change_S_To_Arrow(board, rowStart, colStart+1)
		}
	}

	// CheckLeft
	if colStart != 0 {
		if board[rowStart][colStart-1] == "E" {
			return board
		} else if board[rowStart][colStart-1] == "S" {
			board[rowStart][colStart] = "⬅️"
			Change_S_To_Arrow(board, rowStart, colStart-1)
		}
	}

	return board

}

// PrintBoard
func PrintBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j] + "  ")
		}
		fmt.Print("\n\n")
	}
}
