package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	total := 25 // nombre total de questions

	fmt.Println("Bienvenue dans MathQuizzer ! 🔢")
	fmt.Println("------------------------------")
	fmt.Println("Tu vas passer par 5 étapes de plus en plus dures :")
	fmt.Println("1 à 5     → Additions/Soustractions simples")
	fmt.Println("6 à 10    → Additions/Soustractions complexes")
	fmt.Println("11 à 15   → Multiplications / Divisions")
	fmt.Println("16 à 20   → Équations")
	fmt.Println("21 à 25   → Matrices 🧠")
	fmt.Println()

	for i := 1; i <= total; i++ {
		var step int

		switch {
		case i >= 1 && i <= 5:
			step = 1
		case i >= 6 && i <= 10:
			step = 2
		case i >= 11 && i <= 15:
			step = 3
		case i >= 16 && i <= 20:
			step = 4
		case i >= 21 && i <= 25:
			step = 5
		}

		fmt.Printf("Étape %d - Question %d/%d", step, i, total)

		ex := GenerateExercise(step)

		fmt.Println(ex.Question)
		start := time.Now()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		isMatrix := step == 5

		if CheckAnswer(input, ex, isMatrix) {
			fmt.Printf("✅ Correct ! (%.2fs)\n", time.Since(start).Seconds())
			score++
		} else {
			if isMatrix {
				fmt.Printf("❌ Mauvais ! La bonne réponse était : %s\n", ex.AnswerStr)
			} else {
				fmt.Printf("❌ Mauvais ! La bonne réponse était %.0f\n", ex.Answer)
			}
		}
		fmt.Println()
	}

	fmt.Printf("🎉 Quiz terminé ! Score final : %d/%d\n", score, total)
}
