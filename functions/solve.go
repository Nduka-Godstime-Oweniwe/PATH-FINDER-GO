package PathFinder

// package main

// Creates The Board
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

// PrintBoard
func PrintBoard(board [][]string) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			fmt.Print(board[i][j] + "  ")
		}
		fmt.Print("\n\n")
	}
}

func CopyBoard(board [][]string) [][]string {
	copyBoard := make([][]string, len(board))

	for i := range board {
		copyBoard[i] = make([]string, len(board[i]))
		copy(copyBoard[i], board[i])
	}

	return copyBoard
}

func FindS(board [][]string) (int, int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "S" {
				return i, j
			}
		}
	}
	return 0, 0
}

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

func variations() [][]int {
	final := [][]int{}
	slice := []int{}
	for i := 1; i <= 4; i++ {
		for j := 1; j <= 4; j++ {
			if i != j {
				for k := 1; k <= 4; k++ {
					if k != j && k != i {
						for L := 1; L <= 4; L++ {
							if L != j && L != i && L != k {
								slice = append(slice, i)
								slice = append(slice, j)
								slice = append(slice, k)
								slice = append(slice, L)
							}
						}
						final = append(final, slice)
						slice = []int{}
					}
				}
			}
		}

	}
	return final
}

func functions() ([][]func([][]string, int, int) (bool, bool, bool), [][]string) {
	finalStr := [][]string{}
	str := []string{}
	permutations := variations()
	final := [][]func([][]string, int, int) (bool, bool, bool){}
	function := []func([][]string, int, int) (bool, bool, bool){}
	for i := 0; i < len(permutations); i++ {
		for j := 0; j < 4; j++ {
			if permutations[i][j] == 1 {
				function = append(function, CheckRight)
				str = append(str, "CheckRight")
			} else if permutations[i][j] == 2 {
				function = append(function, CheckLeft)
				str = append(str, "CheckLeft")
			} else if permutations[i][j] == 3 {
				function = append(function, CheckUp)
				str = append(str, "CheckUp")
			} else if permutations[i][j] == 4 {
				function = append(function, CheckDown)
				str = append(str, "CheckDown")
			}
		}
		final = append(final, function)
		finalStr = append(finalStr, str)
		function = []func([][]string, int, int) (bool, bool, bool){}
		str = []string{}
	}
	return final, finalStr
}

func RowAndColAddValue(str string) (int, int) {
	if str == "CheckRight" {
		return 0, 1
	} else if str == "CheckLeft" {
		return 0, -1
	} else if str == "CheckUp" {
		return -1, 0
	} else {
		return 1, 0
	}
}

func Solve(board [][]string, rowS int, colS int, priorities []func([][]string, int, int) (bool, bool, bool), priorString []string) ([][]string, bool) {
	// check right
	Evalue, ValidPos, _ := priorities[0](board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		r, c := RowAndColAddValue(priorString[0])
		board[rowS+r][colS+c] = "S"
		_, solved := Solve(board, rowS+r, colS+c, priorities, priorString)
		if solved {
			return board, true
		}
		board[rowS+r][colS+c] = "."
	}
	// check down
	Evalue, ValidPos, _ = priorities[1](board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		r, c := RowAndColAddValue(priorString[1])
		board[rowS+r][colS+c] = "S"
		_, solved := Solve(board, rowS+r, colS+c, priorities, priorString)
		if solved {
			return board, true
		}
		board[rowS+r][colS+c] = "."
	}

	// check left
	Evalue, ValidPos, _ = priorities[2](board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		r, c := RowAndColAddValue(priorString[2])
		board[rowS+r][colS+c] = "S"
		_, solved := Solve(board, rowS+r, colS+c, priorities, priorString)
		if solved {
			return board, true
		}
		board[rowS+r][colS+c] = "."
	}

	// check up
	Evalue, ValidPos, _ = priorities[3](board, rowS, colS)
	if Evalue {
		return board, true
	}
	if ValidPos {
		r, c := RowAndColAddValue(priorString[3])
		board[rowS+r][colS+c] = "S"
		_, solved := Solve(board, rowS+r, colS+c, priorities, priorString)
		if solved {
			return board, true
		}
		board[rowS+r][colS+c] = "."
	}

	return board, false

}

func sameBoard(a, b [][]string) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}

		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}

	return true
}

func IsContained(AllPaths [][][]string, SinglePath [][]string) bool {
	for i := 0; i < len(AllPaths); i++ {
		if sameBoard(AllPaths[i], SinglePath) {
			return true
		}
	}
	return false
}

func AllValidPaths(board [][]string) [][][]string {
	rowS, colS := FindS(board)
	answer := [][][]string{}
	priorities, priorsString := functions()

	for i := 0; i < 24; i++ {
		copy := CopyBoard(board)
		ans, solved := Solve(copy, rowS, colS, priorities[i], priorsString[i])
		if solved {
			if !IsContained(answer, ans) {
				answer = append(answer, ans)
			}
		} else {
			return answer
		}

	}
	return answer
}

func CountS(board [][]string) int {
	count := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == "S" {
				count++
			}
		}
	}
	return count
}

func SliceOfS(AllPaths [][][]string) []int {
	ans := []int{}
	for i := 0; i < len(AllPaths); i++ {
		ans = append(ans, CountS(AllPaths[i]))

	}
	return ans
}

func SortValidPath(AllPaths [][][]string, numOfmovements []int) {
	for i := 0; i < len(AllPaths); i++ {
		for j := 0; j < len(AllPaths)-1-i; j++ {
			if numOfmovements[j] > numOfmovements[j+1] {
				numOfmovements[j], numOfmovements[j+1] = numOfmovements[j+1], numOfmovements[j]
				AllPaths[j], AllPaths[j+1] = AllPaths[j+1], AllPaths[j]
			}
		}
	}

}

func SolvePathFinder(board [][]string) [][][]string {
	all_paths := AllValidPaths(board)
	SortValidPath(all_paths, SliceOfS(all_paths))
	return all_paths

}
