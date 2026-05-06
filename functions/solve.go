package PathFinder

func Board(slice []string) [][]string {
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

// func PrintBoard(board [][]string) {
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			fmt.Print(board[i][j] + "  ")
// 		}
// 		fmt.Print("\n\n")
// 	}
// }
// func FindS(board [][]string) (int, int) {
// 	for i := 0; i < len(board); i++ {
// 		for j := 0; j < len(board[i]); j++ {
// 			if board[i][j] == "S" {
// 				return i, j
// 			}
// 		}
// 	}
// 	return 0, 0
// }

func CheckRight(board [][]string, rowS, colS int) (bool, bool, bool) {
	if colS != len(board[0])-1 {
		if board[rowS][colS+1] == "E" {
			return true, true, false
		} else if board[rowS][colS+1] == "X" {
			return false, false, false
		} else if board[rowS][colS+1] == "." {
			return false, true, false
		} else {
			return false, false, false
		}
	}
	return false, false, true
}

func CheckLeft(board [][]string, rowS, colS int) (bool, bool, bool) {
	if colS != 0 {
		if board[rowS][colS-1] == "E" {
			return true, true, false
		} else if board[rowS][colS-1] == "X" {
			return false, false, false
		} else if board[rowS][colS-1] == "." {
			return false, true, false
		} else {
			return false, false, false
		}
	}
	return false, false, true
}

func CheckDown(board [][]string, rowS, colS int) (bool, bool, bool) {
	if rowS != len(board)-1 {
		if board[rowS+1][colS] == "E" {
			return true, true, false
		} else if board[rowS+1][colS] == "X" {
			return false, false, false
		} else if board[rowS+1][colS] == "." {
			return false, true, false
		} else {
			return false, false, false
		}
	}
	return false, false, true
}

func CheckUp(board [][]string, rowS, colS int) (bool, bool, bool) {
	if rowS != 0 {
		if board[rowS-1][colS] == "E" {
			return true, true, false
		} else if board[rowS-1][colS] == "X" {
			return false, false, false
		} else if board[rowS-1][colS] == "." {
			return false, true, false
		} else {
			return false, false, false
		}
	}
	return false, false, true
}

func Solve(board [][]string, rowS, colS int) ([][]string, bool) {
	// check right
	Evalue, ValidPos, _ := CheckRight(board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		board[rowS][colS+1] = "➡️"
		_, solved := Solve(board, rowS, colS+1)
		if solved {
			return board, true
		}
		board[rowS][colS+1] = "."
	}
	// check down
	Evalue, ValidPos, _ = CheckDown(board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		board[rowS+1][colS] = "⬇️"
		_, solved := Solve(board, rowS+1, colS)
		if solved {
			return board, true
		}
		board[rowS+1][colS] = "."
	}

	// check left
	Evalue, ValidPos, _ = CheckLeft(board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		board[rowS][colS-1] = "⬅️"
		_, solved := Solve(board, rowS, colS-1)
		if solved {
			return board, true
		}
		board[rowS][colS-1] = "."
	}

	// check up
	Evalue, ValidPos, _ = CheckUp(board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		board[rowS-1][colS] = "⬆️"
		_, solved := Solve(board, rowS-1, colS)
		if solved {
			return board, true
		}
		board[rowS-1][colS] = "."
	}

	return board, false

}

// func main() {
// 	// slice := []string{
// 	// 	"..S..X",
// 	// 	"X...X.",
// 	// 	"E..X..",
// 	// }
// 	// slice := []string{
// 	// 	"......",
// 	// 	"SXEXX.",
// 	// 	".X....",
// 	// }
// 	slice := []string{
// 		"S......................................................",
// 		".......................................................",
// 		".......................................................",
// 		".......................................................",
// 		"......................................................E",
// 	}
// 	board := Board(slice)
// 	rowS, colS := FindS(board)
// 	PrintBoard(board)
// 	final, solved := Solve(board, rowS, colS)
// 	fmt.Println(solved)
// 	PrintBoard(final)

// }
