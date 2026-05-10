package PathFinder

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func GetUserName() string {

	fmt.Println("================ 🔎 The Pathfinder 🔍 ================")
	fmt.Println("Welcome to the Pathfinder🔎")
	var username string

	reader := bufio.NewReader(os.Stdin)

	for {

		// Collect & validate username
		fmt.Print("Please input your username to begin your quest to riches💰💵\nUsername: ")

		name, _ := reader.ReadString('\n')
		username = strings.TrimSpace(name)

		if username == "" {
			fmt.Println("As how na?🤦 You can't be nameless now")
			continue
		}
		if strings.Contains(username, " ") {
			fmt.Println("My G,😎 abeg don't use space. Hint: Use '_' instead")
			continue
		}
		if len(username) > 15 {
			fmt.Println("Nawa for you o,😉 Your name is longer than Methuselah age 🤣. Hint: max. length is 15")
			continue
		}
		fmt.Printf("Hello!👋 %s, The Beast!🫡💯\n", username)
		time.Sleep(3 * time.Second)
		break
	}
	return username
}
