package main

type outcome byte

const (
	lose outcome = 0
	draw         = 3
	win          = 6
)

func getOutcome(outcome string) outcome {
	switch outcome {
	case "X":
		return lose
	case "Y":
		return draw
	case "Z":
		return win
	default:
		return 0
	}
}
