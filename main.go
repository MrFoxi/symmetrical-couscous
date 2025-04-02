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
	http.HandleFunc("/", serveQuiz)
	http.HandleFunc("/check", checkAnswer)

	fmt.Println("🌐 Serveur lancé sur http://localhost:8080")
	http.ListenAndServe(":8080", nil)
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
	questionNumber++
	if questionNumber > totalQuestions {
		final := fmt.Sprintf("Quiz terminé ! Score final : %d/%d", score, totalQuestions)
		tmpl := `<html><head><title>Fin</title></head><body><h2>{{.}}</h2><a href='/restart'>Recommencer</a></body></html>`
		t := template.Must(template.New("final").Parse(tmpl))
		t.Execute(w, final)
		return
	}

	step := getStep(questionNumber)
	currentExercise = GenerateExercise(step)

	tmpl := `<html><head><title>MathQuizzer Web</title></head><body>
		<h3>Question {{.Number}}/{{.Total}}</h3>
		<h2>{{.Question}}</h2>
		<form action='/check' method='POST'>
			<input type='text' name='answer' autofocus />
			<button type='submit'>Vérifier</button>
		</form>
	</body></html>`

	t := template.Must(template.New("quiz").Parse(tmpl))
	t.Execute(w, struct {
		Question string
		Number   int
		Total    int
	}{
		Question: currentExercise.Question,
		Number:   questionNumber,
		Total:    totalQuestions,
	})
}

func checkAnswer(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userInput := r.FormValue("answer")
	ok := CheckAnswer(userInput, currentExercise, currentExercise.AnswerStr != "")

	if ok {
		score++
	}

	feedback := ""
	if ok {
		feedback = "✅ Bonne réponse !"
	} else {
		if currentExercise.AnswerStr != "" {
			feedback = fmt.Sprintf("❌ Mauvais ! La bonne réponse était : %s", currentExercise.AnswerStr)
		} else {
			feedback = fmt.Sprintf("❌ Mauvais ! La bonne réponse était : %.0f", currentExercise.Answer)
		}
	}

	tmpl := `<html><head><title>Résultat</title></head><body>
		<p>{{.}}</p>
		<a href='/'>Question suivante</a>
	</body></html>`
	t := template.Must(template.New("result").Parse(tmpl))
	t.Execute(w, feedback)
}

func restartQuiz(w http.ResponseWriter, r *http.Request) {
	questionNumber = 0
	score = 0
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func init() {
	http.HandleFunc("/restart", restartQuiz)
}
