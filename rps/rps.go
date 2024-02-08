package rps

import (
	"math/rand"
	"strconv"
)

const (
	ROCK     = 0 //piedra.vence a las tijeras.(tijeras +1) % 3 = 0
	PAPER    = 1 //papel.vence a la piedra.(piedra +1) % 3 = 1
	SCISSORS = 2 //tijeras.vence a papel.(papel +1) % 3 = 2
)

type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// mensaje para cuando gana
var winMessages = []string{
	"¡Enhorabuena, has ganado!",
	"¡Fantástico, has ganado!",
	"¡Excelente, has ganado!",
}

// mensaje para cuando pierde
var loseMessages = []string{
	"¡Que lastima, has perdido!",
	"¡Intentalo de nuevo!",
	"¡No te rindas, has perdido!",
}

// mensaje para cuando empata
var drawMessages = []string{
	"¡Hoy es un día de empate!",
	"¡Hoy es un día de empate!",
	"¡Hoy es un día de empate!",
}

// variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string
	var computerChoiceInt int

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió Piedra"
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió Papel"
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió Tijeras"
	}

	messageInt := rand.Intn(3)
	var message string
	if playerValue == computerValue {
		roundResult = "EMPATE"
		message = drawMessages[messageInt]
		
	} else if playerValue == (computerValue+1)%3{
		PlayerScore++
		roundResult = "GANASTE"
		message = winMessages[messageInt]
		
	} else {
		ComputerScore++
		message = loseMessages[messageInt]
		roundResult = "PERDISTE , LA COMPUTADORA GANA"
		
	}


	return Round{
		Message:           message,
		ComputerChoice:    computerChoice,
		RoundResult:       roundResult,
		ComputerChoiceInt: computerChoiceInt,
		ComputerScore:     strconv.Itoa(ComputerScore),
		PlayerScore:       strconv.Itoa(PlayerScore),
	}
}