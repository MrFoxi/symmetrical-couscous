package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var (
	currentExercise Exercise
	questionNumber  int
	score           int
	totalQuestions  = 25
)

func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/quiz", serveQuiz)
	http.HandleFunc("/restart", restartQuiz)

	fmt.Println("🌐 Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/index.html"))
	tmpl.Execute(w, nil)
}

func getStep(question int) int {
	switch {
	case question >= 1 && question <= 5:
		return 1
	case question >= 6 && question <= 10:
		return 2
	case question >= 11 && question <= 15:
		return 3
	case question >= 16 && question <= 20:
		return 4
	case question >= 21 && question <= 25:
		return 5
	default:
		return 1
	}
}

func serveQuiz(w http.ResponseWriter, r *http.Request) {
	check := CheckAnswer(r.FormValue("answer"), currentExercise, currentExercise.AnswerStr != "")
	if check && questionNumber > 0 {
		score++
	}

	if questionNumber >= totalQuestions {
		finalScore := fmt.Sprintf("%d/%d", score, totalQuestions)
		tmpl := template.Must(template.ParseFiles("html/end.html")) // Utilisez le nouveau template
		tmpl.Execute(w, finalScore)
		return
	}

	notification := getNotification(check, questionNumber, score, totalQuestions)
	questionNumber++

	step := getStep(questionNumber)
	currentExercise = GenerateExercise(step)
	tmpl := template.Must(template.ParseFiles("html/quiz.html"))
	tmpl.Execute(w, struct {
		Question     string
		Number       int
		Total        int
		Notification string
	}{
		Question:     currentExercise.Question,
		Number:       questionNumber,
		Total:        totalQuestions,
		Notification: notification,
	})
}

func getNotification(check bool, questionNumber int, score int, totalQuestions int) string {
	if questionNumber == 0 {
		return "Bienvenue dans le quiz !"
	} else if check {
		return fmt.Sprintf("Bonne réponse ! Score actuel : %d/%d", score, totalQuestions)
	} else {
		return fmt.Sprintf("Réponse incorrecte. Score actuel : %d/%d", score, totalQuestions)
	}
}

func restartQuiz(w http.ResponseWriter, r *http.Request) {
	questionNumber = 0
	score = 0
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
