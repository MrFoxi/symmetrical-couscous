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
	lastFeedback    string
)

<<<<<<< HEAD:Ancien projet/main.go
// func main() {
// 	http.HandleFunc("/", serveQuiz)
// 	http.HandleFunc("/check", checkAnswer)
// 	http.HandleFunc("/restart", restartQuiz)
=======
func main() {
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/quiz", serveQuiz)
	http.HandleFunc("/restart", restartQuiz)
>>>>>>> 2696e7498776629e580e30fa8f3e03689e6acacf:main.go

// 	fmt.Println("🌐 Serveur lancé sur http://localhost:8080")
// 	http.ListenAndServe(":8080", nil)
// }

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
<<<<<<< HEAD:Ancien projet/main.go
	if questionNumber >= totalQuestions {
		final := fmt.Sprintf("Quiz terminé ! Score final : %d/%d", score, totalQuestions)
		tmpl := `<html><head><title>Fin</title></head><body><h2>{{.}}</h2><a href='/restart'>Recommencer</a></body></html>`
		t := template.Must(template.New("final").Parse(tmpl))
		t.Execute(w, final)
		return
	}

	questionNumber++
	// step := getStep(questionNumber)
	// currentExercise = GenerateExercise(step)

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
		lastFeedback = "✅ Bonne réponse !"
	} else {
		if currentExercise.AnswerStr != "" {
			lastFeedback = fmt.Sprintf("❌ Mauvais ! La bonne réponse était : %s", currentExercise.AnswerStr)
		} else {
			lastFeedback = fmt.Sprintf("❌ Mauvais ! La bonne réponse était : %.0f", currentExercise.Answer)
		}
	}

	tmpl := `<html><head><title>Résultat</title>
		<script>
			setTimeout(function() {
				window.location.href = '/';
			}, 3000);
		</script>
	</head><body>
		<p>{{.}}</p>
		<p>Redirection dans 3 secondes...</p>
	</body></html>`
	t := template.Must(template.New("result").Parse(tmpl))
	t.Execute(w, lastFeedback)
=======
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
>>>>>>> 2696e7498776629e580e30fa8f3e03689e6acacf:main.go
}

func restartQuiz(w http.ResponseWriter, r *http.Request) {
	questionNumber = 0
	score = 0
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
