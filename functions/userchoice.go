package PathFinder

import (
	"fmt"
)

func Input(prompt string) string {
	var answer string
	fmt.Print(prompt)
	fmt.Scan(&answer)
	return answer

}
func Choice() string {
	fmt.Println("1. Shortest Path")
	fmt.Println("2. Longest Path")
	fmt.Println("3. All Path")
	fmt.Println("4. Go Back")
	option := ""
	for option == "" {
		option = Input("Select Option: ")
		if option != "1" && option != "2" && option != "3" && option != "4" {
			fmt.Println("Invalid Option")
			option = ""
		} else {
			break
		}
	}
	return option

}

func UserChoice(board [][]string) (bool, bool) {
	paths := Choice()
	validPaths := SolvePathFinder(board)
	if len(validPaths) == 0 {
		return true, false
	}
	if paths == "1" {
		count := CountS(validPaths[0])
		path := 1
		for i := 0; i < len(validPaths); i++ {
			if count == CountS(validPaths[i]) {
				fmt.Printf("Path %d:\n", path)
				path++
				PrintBoard(validPaths[i])
				fmt.Print("\n\n")
			}
		}

	} else if paths == "2" {
		count := CountS(validPaths[len(validPaths)-1])
		path := 1
		for i := 0; i < len(validPaths); i++ {
			if count == CountS(validPaths[i]) {
				fmt.Printf("Path %d:\n", path)
				path++
				PrintBoard(validPaths[i])
				fmt.Print("\n\n")
			}
		}
	} else if paths == "3" {
		path := 1
		for i := 0; i < len(validPaths); i++ {
			fmt.Printf("Path %d:\n", path)
			path++
			PrintBoard(validPaths[i])
			fmt.Print("\n\n")
		}
	} else {
		return false, true
	}

	Input("Type Anything to continue 👉: ")
	return true, true
}
