package main

import (
	"fmt"
)

func main() {
	// RPS 1 function
	fmt.Println(Rps("rock", "scissors"))     // P1 WINS
	fmt.Println(Rps("rock", "paper"))        // P1 WINS
	fmt.Println(Rps("paper", "scissors"))    // P2 WINS
	fmt.Println(Rps("rock", "paper"))        // P2 WINS
	fmt.Println(Rps("scissors", "scissors")) // DRAW

	// RPS 2 function
	fmt.Println(Rps2("rock", "scissors"))     // P1 WINS
	fmt.Println(Rps2("scissors", "paper"))    // P1 WINS
	fmt.Println(Rps2("paper", "paper"))       // DRAW
	fmt.Println(Rps2("scissors", "scissors")) // DRAW
	fmt.Println(Rps2("paper", "rock"))        // P1 WINS

}

func Rps(p1, p2 string) string { // My Solution

	if p1 == "rock" && p2 == "scissors" {
		return "Player 1 won!"
	}
	if p1 == "scissors" && p2 == "rock" {
		return "Player 2 won!"
	}
	if p1 == "paper" && p2 == "rock" {
		return "Player 1 won!"
	}
	if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	}
	if p1 == "rock" && p2 == "paper" {
		return "Player 2 won!"
	}
	if p1 == "scissors" && p2 == "rock" {
		return "Player 2 won!"
	}
	if p1 == "scissors" && p2 == "paper" {
		return "Player 1 won!"
	}
	if p1 == "paper" && p2 == "scissors" {
		return "Player 2 won!"

	} else {
		return "Draw!"
	}
}

func Rps2(p1, p2 string) string { // Solution I found

	var m = map[string]string{"rock": "paper", "paper": "scissors", "scissors": "rock"}

	if p1 == p2 {
		return "Draw!"
	}
	if m[p1] == p2 {
		return "Player 2 won!"
	}
	return "Player 1 won!"
}
