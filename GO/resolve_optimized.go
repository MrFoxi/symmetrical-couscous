package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"gonum.org/v1/gonum/mat"
)

func main() {
	fmt.Println("[INFO] Chargement des matrices depuis matrices.json...")

	f, err := os.Open("matrices.json")
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture du fichier :", err)
		return
	}
	defer f.Close()

	var data map[string][][]float64
	if err := json.NewDecoder(f).Decode(&data); err != nil {
		fmt.Println("Erreur lors du décodage JSON :", err)
		return
	}

	a := data["A"]
	b := data["B"]

	rows, cols := len(a), len(b[0])
	n := len(a[0]) // nombre de colonnes de A = nombre de lignes de B

	flatA := make([]float64, 0, rows*n)
	for _, row := range a {
		flatA = append(flatA, row...)
	}

	flatB := make([]float64, 0, n*cols)
	for _, row := range b {
		flatB = append(flatB, row...)
	}

	matA := mat.NewDense(rows, n, flatA)
	matB := mat.NewDense(n, cols, flatB)

	fmt.Println("[INFO] Multiplication avec gonum...")
	start := time.Now()
	var result mat.Dense
	result.Mul(matA, matB)
	elapsed := time.Since(start)

	fmt.Printf("[RESULT] Temps de résolution (Go + gonum) : %s\n", elapsed)
}
