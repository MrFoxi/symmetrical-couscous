package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Exercise struct {
	Question string
	Answer   float64
}

func GenerateExercise(level int) Exercise {
	rand.Seed(time.Now().UnixNano())

	switch level {
	case 1:
		return generateEasy()
	case 2:
		return generateMedium()
	case 3:
		return generateHard()
	default:
		return generateEasy()
	}
}

func generateEasy() Exercise {
	a := rand.Intn(10)
	b := rand.Intn(10)
	op := rand.Intn(3)
	var question string
	var answer float64

	switch op {
	case 0:
		question = fmt.Sprintf("Combien font %d + %d ?", a, b)
		answer = float64(a + b)
	case 1:
		question = fmt.Sprintf("Combien font %d - %d ?", a, b)
		answer = float64(a - b)
	case 2:
		question = fmt.Sprintf("Combien font %d * %d ?", a, b)
		answer = float64(a * b)
	}

	return Exercise{Question: question, Answer: answer}
}

func generateMedium() Exercise {
	a := rand.Intn(10) + 1 // éviter 0
	b := rand.Intn(10)
	x := rand.Intn(10)
	c := a*x + b

	question := fmt.Sprintf("Résous : %dx + %d = %d. Quelle est la valeur de x ?", a, b, c)
	answer := float64(x)

	return Exercise{Question: question, Answer: answer}
}

func generateHard() Exercise {
	a := rand.Intn(5) + 1
	b := rand.Intn(3) + 2

	// Exemple : dérivée de ax^b
	question := fmt.Sprintf("Quelle est la dérivée de f(x) = %dx^%d ?", a, b)
	answer := float64(a * b) // ax^b → abx^(b-1), on attend juste le coefficient ici

	return Exercise{Question: question, Answer: answer}
}

func CheckAnswer(userAnswer float64, correctAnswer float64) bool {
	return userAnswer == correctAnswer
}
