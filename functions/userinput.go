package PathFinder

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	fmt.Println(UserInput())
}

func GetUserName() string {
	fmt.Println("=== Welcome To The Pathfinder ===")
	var username string

	reader := bufio.NewReader(os.Stdin)

	for {
		// Collect & validate username
		fmt.Print("Please, input your username to begin your quest to riches\nUsername: ")

		name, _ := reader.ReadString('\n')
		username = strings.TrimSpace(name)

		if username == "" {
			fmt.Println("Error: Username cannot be empty")
			continue
		}
		if strings.Contains(username, " ") {
			fmt.Println("Error: Username cannot contain spaces. Hint: Use '_' instead")
			continue
		}
		if len(username) > 15 {
			fmt.Println("Error: Username is too long. Hint: max. length is 15")
			continue
		}
		fmt.Printf("Hello, %s, The ToloTolo\n", username)
		break
	}
	return username
}

func GetUserChoice() string {
	reader := bufio.NewReader(os.Stdin)
	username := GetUserName()
	fmt.Println("1: Input entire board rows at once")
	fmt.Println("2: Input board row sequentially")
	fmt.Println("3: Exit")
	fmt.Printf("To start your journey %s, select from the options above 👆.\n", username)
	option := ""
	for {
		fmt.Print("Choose option: ")
		option, _ = reader.ReadString('\n')
		option = strings.TrimSpace(option)
		if option != "1" && option != "2" && option != "3" {
			fmt.Println("Invalid Option!")
			continue
		}
		break
	}
	return option
}

func CheckCharacters(str string) bool {
	countS := 0
	countE := 0
	for _, v := range str {
		if v != 'S' && v != 'E' && v != 'X' && v != '.' {
			return false
		} else if v == 'S' {
			countS++
		} else if v == 'E' {
			countE++
		}
	}
	if countS > 1 || countE > 1 {
		return false
	} else {
		return true
	}
}

func CheckE(str string) bool {
	for _, v := range str {
		if v == 'E' {
			return true
		}
	}
	return false
}

func CheckS(str string) bool {
	for _, v := range str {
		if v == 'S' {
			return true
		}
	}
	return false
}

func EachRow(num int) (string, bool, bool) {
	reader := bufio.NewReader(os.Stdin)
	rows := ""

	for {
		fmt.Printf("Enter Row %v or type \"done\" to stop: ", num)

		rows, _ = reader.ReadString('\n')
		rows = strings.TrimSpace(rows)

		if rows == "" {
			fmt.Println("String Cannot Be Empty")
			continue

		} else if strings.ToUpper(rows) == "DONE" {
			break

		} else if !CheckCharacters(rows) {
			fmt.Println("Only 1 \"S\", 1 \"E\" and as many \"X\" and \".\" are allowed ")
			continue
		}

		break
	}

	return rows, CheckS(rows), CheckE(rows)
}

func RowByRow() []string {
	S := 0
	E := 0
	row := ""
	Sbool := true
	Ebool := true
	num := 1
	slice := []string{}

	for {
		row, Sbool, Ebool = EachRow(num)

		if strings.ToUpper(row) == "DONE" {
			break

		} else if (S >= 1 && Sbool) || (E >= 1 && Ebool) {
			fmt.Println("You can only have 1 E and 1 S in the entire path")
			continue

		} else if len(slice) != 0 {
			if len(slice[0]) != len(row) {
				fmt.Println("Oooops! Invalid number of characters. Each row must have equal length")
				continue
			}
		}

		if Sbool {
			S++
		}

		if Ebool {
			E++
		}

		slice = append(slice, row)
		num++
	}

	return slice
}

func Everything() []string {
	reader := bufio.NewReader(os.Stdin)
	answer := ""
	slice := []string{}
	for {
		fmt.Print("Enter All Rows(Each separated by a space): ")
		answer, _ = reader.ReadString('\n')
		answer = strings.TrimSpace(answer)
		if answer == "" {
			fmt.Println("Cannot accept and empty Input")
			continue
		}
		slice = strings.Fields(answer)
		check := true
		S := 0
		E := 0

		for i := 0; i < len(slice); i++ {
			if S >= 1 && CheckS(slice[i]) {
				fmt.Println("Pls Ensure you have only 1 S in the entire board")
				check = false
				break
			}

			if E >= 1 && CheckE(slice[i]) {
				fmt.Println("Pls Ensure you have only 1 E in the entire board")
				check = false
				break
			}

			if !CheckCharacters(slice[i]) {
				fmt.Println("Only 1 \"S\", 1 \"E\" and as many \"X\" and \".\" are allowed per row ")
				check = false
				break

			}
			if len(slice[0]) != len(slice[i]) {
				fmt.Println("Pls Ensure Each Row are of equal length")
				check = false
				break
			}

			if CheckE(slice[i]) {
				E++
			}

			if CheckS(slice[i]) {
				S++
			}
		}
		if check {
			break
		}

	}
	return slice
}

func UserInput() []string {
	userchoice := GetUserChoice()
	if userchoice == "1" {
		return RowByRow()
	} else if userchoice == "2" {
		return Everything()
	} else {
		return []string{}
	}
}
