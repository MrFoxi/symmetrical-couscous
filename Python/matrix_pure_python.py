import json
import time

def multiply_matrices(a, b):
    size = len(a)
    result = [[0.0 for _ in range(size)] for _ in range(size)]
    for i in range(size):
        for j in range(size):
            total = 0.0
            for k in range(size):
                total += a[i][k] * b[k][j]
            result[i][j] = total
    return result

print("[INFO] Chargement des matrices depuis matrices.json...")
with open("matrices.json", "r") as f:
    data = json.load(f)

A = data["A"]
B = data["B"]

print("[INFO] Multiplication en cours (Python pur)...")
start = time.time()
_ = multiply_matrices(A, B)
elapsed = time.time() - start

print(f"[RESULT] Temps de calcul (Python pur) : {elapsed:.4f} secondes")
