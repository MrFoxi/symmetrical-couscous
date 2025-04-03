package main

import (
    "encoding/json"
    "fmt"
    "math/rand"
    "os"
    "time"
)

const size = 1000

func generateMatrix(size int) [][]float64 {
    matrix := make([][]float64, size)
    for i := range matrix {
        matrix[i] = make([]float64, size)
        for j := range matrix[i] {
            matrix[i][j] = rand.Float64() * 100
        }
    }
    return matrix
}

func main() {
    rand.Seed(time.Now().UnixNano())

    fmt.Println("⏳ Génération des matrices...")
    a := generateMatrix(size)
    b := generateMatrix(size)

    fmt.Println("💾 Sauvegarde dans matrices.json")
    f, _ := os.Create("matrices.json")
    json.NewEncoder(f).Encode(map[string][][]float64{
        "A": a,
        "B": b,
    })
    f.Close()

    fmt.Println("✅ Matrices générées avec succès.")
}
