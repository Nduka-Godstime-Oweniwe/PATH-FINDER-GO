package main

import "fmt"

func CreateBoard(rows []string) [][]rune{
	var board [][] rune

	for _, row := range rows {
		board = append(board, [] rune(row))

	}
	 return board
	

}

func main(){
	rows := []string{
		"ABC",
		"DEF",
		"GHI",
	}

	board := CreateBoard(rows)

	for _, row := range board{
		fmt.Println(string(row))
	}
}