package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type Exercise struct {
	Question  string
	Answer    float64
	AnswerStr string
}

// func GenerateExercise(step int) Exercise {
// 	switch step {
// 	case 1:
// 		return generateBasicAddSub(0, 10)
// 	case 2:
// 		return generateBasicAddSub(20, 100)
// 	case 3:
// 		return generateMultDiv()
// 	case 4:
// 		return generateEquation()
// 	case 5:
// 		// return generateMatrix()
// 	default:
// 		return generateBasicAddSub(0, 10)
// 	}
// }

func generateBasicAddSub(min int, max int) Exercise {
	a := rand.Intn(max-min) + min
	b := rand.Intn(max-min) + min
	op := rand.Intn(2)
	var question string
	var answer float64

	if op == 0 {
		question = fmt.Sprintf("\nCombien font %d + %d ?", a, b)
		answer = float64(a + b)
	} else {
		question = fmt.Sprintf("\nCombien font %d - %d ?", a, b)
		answer = float64(a - b)
	}

	return Exercise{Question: question, Answer: answer}
}

func generateMultDiv() Exercise {
	a := rand.Intn(12) + 1
	b := rand.Intn(12) + 1
	op := rand.Intn(2)
	var question string
	var answer float64

	if op == 0 {
		question = fmt.Sprintf("\nCombien font %d * %d ?", a, b)
		answer = float64(a * b)
	} else {
		question = fmt.Sprintf("\nCombien font %d / %d ?", a*b, a)
		answer = float64(b)
	}

	return Exercise{Question: question, Answer: answer}
}

func generateEquation() Exercise {
	a := rand.Intn(9) + 1
	b := rand.Intn(10)
	x := rand.Intn(10)
	c := a*x + b

	question := fmt.Sprintf("Résous : %dx + %d = %d. Quelle est la valeur de x ?", a, b, c)
	answer := float64(x)

	return Exercise{Question: question, Answer: answer}
}

// func generateMatrix() Exercise {
// 	a := [2][2]int{{rand.Intn(10), rand.Intn(10)}, {rand.Intn(10), rand.Intn(10)}}
// 	b := [2][2]int{{rand.Intn(10), rand.Intn(10)}, {rand.Intn(10), rand.Intn(10)}}
// 	var result [2][2]int

// 	for i := 0; i < 2; i++ {
// 		for j := 0; j < 2; j++ {
// 			result[i][j] = a[i][j] + b[i][j]
// 		}
// 	}

// 	formattedAnswer := fmt.Sprintf("[[%d %d] [%d %d]]", result[0][0], result[0][1], result[1][0], result[1][1])
// 	question := fmt.Sprintf("Additionne ces matrices 2x2 :\nA = %v\nB = %v\nRéponds sous le format : [[a b] [c d]]", a, b)

// 	return Exercise{Question: question, AnswerStr: formattedAnswer}
// }

func CheckAnswer(userInput string, ex Exercise, isMatrix bool) bool {
	if isMatrix {
		return strings.TrimSpace(userInput) == strings.TrimSpace(ex.AnswerStr)
	}
	userAnswer, err := strconv.ParseFloat(strings.TrimSpace(userInput), 64)
	if err != nil {
		return false
	}
	return userAnswer == ex.Answer
}
