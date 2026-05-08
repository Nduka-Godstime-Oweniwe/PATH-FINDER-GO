package main

import (
	"strings"
	// "fmt"
)

func UserInput(str []string) []string {
	if len(str) == 0 {
		return []string{}
	}

	// Dealing with the characters in each row and length of row
	countrowlength := 0
	for i := 0; i < len(str[0]); i++ {
		countrowlength++
	}

	stringArray := []string{}
	wholestring := ""
	for j := 0; j < len(str); j++ {
		row := str[j]
		if len(row) != countrowlength {
			return []string{}
		}
		newstring := ""
		for i := 0; i < len(row); i++ {
			if strings.Contains(row, "S") == true {
				count := strings.Count(row, "S")
				if count != 1 {
					return []string{}
				}
			}
			if strings.Contains(row, "E") == true {
				count := strings.Count(row, "E")
				if count != 1 {
					return []string{}
				}
			}
			if row[i] != 'S' && row[i] != 'E' && row[i] != 'X' && row[i] != '.' {
				return []string{}
			}
			newstring += string(row[i])
		}
		wholestring += newstring
		stringArray = append(stringArray, newstring)
	}

	countS := strings.Count(wholestring, "S")
	if strings.Contains(wholestring, "S") && countS != 1 {
		return []string{}
	}
	countE := strings.Count(wholestring, "E")
	if strings.Contains(wholestring, "E") && countE != 1 {
		return []string{}
	}
	if strings.Contains(wholestring, "E") == false || strings.Contains(wholestring, "S") == false {
		return []string{}
	}

	return stringArray
}

// func main () {
// 	a:= []string{"...S..X.XX.", "......XEXX.", "......X.XXE", "......X.XX."} // duplicate E
// 	b:= []string{"......X.XX.", "......X.XX.", "......X.XXE", "......X.XX."} // Absent S
// 	c:= []string{"....0...XX.", "......X.XX.", "......X.XXE", "......X.XX."} // invalid character 0
// 	d:= []string{"...S..X.XX.", "......a.XX.", "......X.XXE", "......X.XX."} // invalid character a
// 	e:= []string{"...S..X.XX.", "...,..X.XX.", "......X.XXE", "......X.XX."} // invalid character ,
// 	f:= []string{"...S..X.XX.", ".....X..XX.", "......X.XX.", "......X.XX."} // absent E
// 	g:= []string{"......X.XX.", ".....X..XX.", "......X.XXE", "...s..X.XX."} // lowercase character s
// 	h:= []string{"...S..X.XX.", ".....X.XX.", "......X.XXE", "...s..X.XX."} // unequal row
// 	i:= []string{"...S..X.XX.", "......X.XX.", "......X.XXE", "...s....X.XX."} // unequal row
// 	j:= []string{"...S..X.XX.", "......XSXX.", "......X.XXE", "......X.XX."} // duplicate S
// 	k:= []string{"...S..X.XX.", "......XEXX.", "......X.XX.", "......X.XX."} // valid string array
// 	fmt.Printf("%#v\n", UserInput(a))
// 	fmt.Printf("%#v\n", UserInput(b))
// 	fmt.Printf("%#v\n", UserInput(c))
// 	fmt.Printf("%#v\n", UserInput(d))
// 	fmt.Printf("%#v\n", UserInput(e))
// 	fmt.Printf("%#v\n", UserInput(f))
// 	fmt.Printf("%#v\n", UserInput(g))
// 	fmt.Printf("%#v\n", UserInput(h))
// 	fmt.Printf("%#v\n", UserInput(i))
// 	fmt.Printf("%#v\n", UserInput(j))
// 	fmt.Printf("%#v\n", UserInput(k))
// }
