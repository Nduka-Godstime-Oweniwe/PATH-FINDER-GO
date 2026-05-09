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
	option := ""
	for option == "" {
		option = Input("Select Option: ")
		if option != "1" && option != "2" && option != "3" {
			fmt.Println("Invalid Option")
			option = ""
		} else {
			break
		}
	}
	return option

}

func UserReplay(board [][]string) {
	paths := Choice()
	validPaths := SolvePathFinder(board)
	if paths == "1" {
		count := CountS(validPaths[0])
		paths := 1
		for i := 0; i < len(validPaths); i++ {
			if count == CountS(validPaths[i]) {
				fmt.Printf("Path %d:\n", paths)
				paths++
				PrintBoard(validPaths[i])
				fmt.Print("\n\n")
			}
		}

	} else if paths == "2" {
		count := CountS(validPaths[len(validPaths)-1])
		paths := 1
		for i := 0; i < len(validPaths); i++ {
			if count == CountS(validPaths[i]) {
				fmt.Printf("Path %d:\n", paths)
				paths++
				PrintBoard(validPaths[i])
				fmt.Print("\n\n")
			}
		}
	} else {
		paths := 1
		for i := 0; i < len(validPaths); i++ {
			fmt.Printf("Path %d:\n", paths)
			paths++
			PrintBoard(validPaths[i])
			fmt.Print("\n\n")
		}
	}
}
