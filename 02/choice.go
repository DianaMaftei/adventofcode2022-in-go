package main

type choice byte

const (
	rock     choice = 1
	paper           = 2
	scissors        = 3
)

func getChoice(choice string) choice {
	switch choice {
	case "A":
		return rock
	case "B":
		return paper
	case "C":
		return scissors
	default:
		return 0
	}
}
