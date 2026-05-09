package PathFinder

import "fmt"

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
