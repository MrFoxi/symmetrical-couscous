package main

import (
    "encoding/json"
    "fmt"
    "os"
    "time"
)

const size = 1000

func multiplyMatrices(a, b [][]float64) [][]float64 {
    result := make([][]float64, size)
    for i := range result {
        result[i] = make([]float64, size)
        for j := 0; j < size; j++ {
            sum := 0.0
            for k := 0; k < size; k++ {
                sum += a[i][k] * b[k][j]
            }
            result[i][j] = sum
        }
    }
    return result
}

func main() {
    fmt.Println("📂 Chargement de matrices.json...")
    f, _ := os.Open("matrices.json")
    defer f.Close()

    var data map[string][][]float64
    json.NewDecoder(f).Decode(&data)

    a := data["A"]
    b := data["B"]

    fmt.Println("🚀 Résolution Go...")
    start := time.Now()
    _ = multiplyMatrices(a, b)
    elapsed := time.Since(start)

    fmt.Printf("✅ Temps de résolution (Go) : %s\n", elapsed)
}
