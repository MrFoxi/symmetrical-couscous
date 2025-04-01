package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	score := 0
	total := 5 // nombre de questions

	fmt.Println("Bienvenue dans MathQuizzer ! 🔢")
	fmt.Println("------------------------------")

	fmt.Print("Choisis ton niveau (1 = Facile, 2 = Moyen, 3 = Difficile) : ")
	levelInput, _ := reader.ReadString('\n')
	levelInput = strings.TrimSpace(levelInput)
	level, err := strconv.Atoi(levelInput)
	if err != nil || level < 1 || level > 3 {
		fmt.Println("Niveau invalide, niveau facile sélectionné par défaut.")
		level = 1
	}

	for i := 1; i <= total; i++ {
		ex := GenerateExercise(level)
		fmt.Printf("Question %d/%d: %s\n", i, total, ex.Question)

		start := time.Now()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		userAnswer, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("⛔ Entrée invalide. Réponse considérée comme fausse.")
			continue
		}

		duration := time.Since(start).Seconds()

		if CheckAnswer(userAnswer, ex.Answer) {
			fmt.Printf("✅ Correct ! (%.2fs)\n", duration)
			score++
		} else {
			fmt.Printf("❌ Mauvais ! La bonne réponse était %.0f\n", ex.Answer)
		}
		fmt.Println()
	}

	fmt.Printf("🎉 Quiz terminé ! Score final : %d/%d\n", score, total)
}
