package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

/* I/O Functions */

func getPlayerNames() (string, string) {
	p1, _ := readInput("What's the player 1 name? ")
	p2, _ := readInput("What's the player 2 name? ")

	return p1, p2
}

func readInput(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)
	text, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return text, err
}

func printMatrix(matrix [][]byte) {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			fmt.Print(string(matrix[i][j]))
		}
		fmt.Print("\n")
	}
}

/* Vitory/Draw conditions*/

func checkIfFullMatrix(matrix [][]byte) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == ' ' {
				return false
			}
		}
	}
	return true
}

func checkIfSomeoneWonColumns(matrix [][]byte) bool {
	return matrix[0][0] != ' ' && matrix[0][0] == matrix[1][0] && matrix[1][0] == matrix[2][0] || matrix[0][1] != ' ' && matrix[0][1] == matrix[1][1] && matrix[1][1] == matrix[2][1] || matrix[0][2] != ' ' && matrix[0][2] == matrix[1][2] && matrix[1][2] == matrix[2][2]
}

func checkIfSomeoneWonRow(matrix [][]byte) bool {
	return matrix[0][0] != ' ' && matrix[0][0] == matrix[0][1] && matrix[0][1] == matrix[0][2] || matrix[1][0] != ' ' && matrix[1][0] == matrix[1][1] && matrix[1][1] == matrix[1][2] || matrix[2][0] != ' ' && matrix[2][0] == matrix[2][1] && matrix[2][1] == matrix[2][2]
}

func checkIfSomeoneWonDiagonal(matrix [][]byte) bool {
	return matrix[0][0] != ' ' && matrix[0][0] == matrix[1][1] && matrix[1][1] == matrix[2][2] || matrix[0][2] != ' ' && matrix[0][2] == matrix[1][1] && matrix[1][1] == matrix[2][0]
}

func checkIfWon(matrix [][]byte) bool {
	return checkIfSomeoneWonDiagonal(matrix) || checkIfSomeoneWonRow(matrix) || checkIfSomeoneWonColumns(matrix) || checkIfFullMatrix(matrix)
}

func checkIfFinished(matrix [][]byte) bool {
	return checkIfFullMatrix(matrix)
}

/* Play functions */

func doMove(matrix [][]byte, i int, j int, player byte) bool {

	if matrix[i][j] != ' ' {
		return false
	}

	matrix[i][j] = player
	return true
}

func getMovePos(player int) (int, int) {

	stri := "Player " + strconv.Itoa(player) + " turn, what i pos do you want to play? (number from 1 to 3) "
	strj := "Player " + strconv.Itoa(player) + " turn, what j pos do you want to play? (number from 1 to 3) "

	var err error
	var i, j int64

	fmt.Println(stri)
	for ok := true; ok; ok = (err == nil) {
		_, err = fmt.Scanf("%d", &i)
	}

	fmt.Println(strj)
	for ok := true; ok; ok = (err == nil) {

		_, err = fmt.Scanf("%d", &j)
	}

	//fmt.Println(i, j)

	return int(i), int(j)

}

func play(matrix [][]byte, move int) (int, bool, bool) {

	if checkIfWon(matrix) {
		return 0, false, true
	}
	if checkIfFinished(matrix) {
		return 0, true, true
	}

	var player1, player2 byte
	player1 = 'X'
	player2 = 'O'
	test := true
	if move%2 == 0 {
		fmt.Println("It's player 1 turn, choose wisely!")
		for ok := true; ok; ok = (test == false) {
			movei, movej := getMovePos(1)
			test = doMove(matrix, movei, movej, player1)
		}
	} else {
		fmt.Println("It's player 2 turn, choose wisely!")

		for ok := true; ok; ok = (test == false) {
			movei, movej := getMovePos(2)
			test = doMove(matrix, movei, movej, player2)
		}
	}

	printMatrix(matrix)

	return move, false, false

}

/* Main "function" */

func main() {
	rand.Seed(523131231230)
	matrix := [][]byte{{' ', ' ', ' '}, {' ', ' ', ' '}, {' ', ' ', ' '}}
	//p1, p2 := getPlayerNames()

	move := rand.Intn(2)
	//fmt.Print(p1, p2)

	won, tie := false, false

	for tie == false && won == false {
		move, tie, won = play(matrix, move)
		move++
	}

	if won {
		fmt.Printf("Player %d won \n", (move-1%2)+1)
	} else {
		fmt.Println("Tied game sorry :( ")
	}

	printMatrix(matrix)
}
