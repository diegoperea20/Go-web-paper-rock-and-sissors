package rps

import (
	"fmt"
	"testing"
)
func TestPlayRound(t *testing.T) {
	for i:= 0; i < 3; i++ {
		round := PlayRound(i)

		switch i{
		case 0:
			fmt.Println("el jugador eligio piedra")
		case 1:
			fmt.Println("el jugador eligio papel")
		case 2:
			fmt.Println("el jugador eligio tijeras")
		}
		fmt.Printf("Mesage: %s\n", round.Message)
		fmt.Printf("Computer Choice: %s\n", round.ComputerChoice)
		fmt.Printf("Round Result: %s\n", round.RoundResult)
		fmt.Printf("Computer Choice Int: %d\n", round.ComputerChoiceInt)
		fmt.Printf("Computer Score: %s\n", round.ComputerScore)
		fmt.Printf("Player Score: %s\n", round.PlayerScore)
		
		fmt.Println("\n========================================")

		
	}
	
}