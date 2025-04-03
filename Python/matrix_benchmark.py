import json
import time
import numpy as np

print("[INFO] Chargement des matrices depuis matrices.json...")
with open("matrices.json", "r") as f:
    data = json.load(f)

A = np.array(data["A"])
B = np.array(data["B"])

print("[INFO] Multiplication en cours (NumPy)...")
start = time.time()
result = np.dot(A, B)
elapsed = time.time() - start

print(f"[RESULT] Temps de calcul (Python + NumPy) : {elapsed:.4f} secondes")
