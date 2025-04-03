import subprocess
import time
import threading
import os

results = {
    "Go (natif)": {"time": 0, "output": ""},
    "Go (gonum)": {"time": 0, "output": ""},
    "Python (NumPy)": {"time": 0, "output": ""},
    "Python (pur)": {"time": 0, "output": ""},
    "JavaScript": {"time": 0, "output": ""},
    "JavaScript (mathjs)": {"time": 0, "output": ""},
    "C": {"time": 0, "output": ""}
}

def run_process(label, command, storage):
    try:
        start = time.time()
        proc = subprocess.run(command, stdout=subprocess.PIPE, stderr=subprocess.PIPE, text=True)
        storage["time"] = time.time() - start
        storage["output"] = proc.stdout.strip() + "\n" + proc.stderr.strip()
    except Exception as e:
        storage["output"] = f"⚠️ Erreur lors de l'exécution de {label} : {str(e)}"
        storage["time"] = -1

gonum_exe = "go/resolve_gonum.exe" if os.name == "nt" else "./go/resolve_gonum"
c_exe = "c/resolve_c.exe" if os.name == "nt" else "./c/resolve_c"

threads = [
    threading.Thread(target=run_process, args=("Go (natif)", ["go", "run", "go/resolve.go"], results["Go (natif)"])),
    threading.Thread(target=run_process, args=("Go (gonum)", [gonum_exe], results["Go (gonum)"])),
    threading.Thread(target=run_process, args=("Python (NumPy)", ["python", "python/matrix_benchmark.py"], results["Python (NumPy)"])),
    # threading.Thread(target=run_process, args=("Python (pur)", ["python", "python/matrix_pure_python.py"], results["Python (pur)"])), 
    threading.Thread(target=run_process, args=("JavaScript", ["node", "js/resolve.js"], results["JavaScript"])),
    threading.Thread(target=run_process, args=("JavaScript (mathjs)", ["node", "js/resolve_lib.js"], results["JavaScript (mathjs)"])),
    threading.Thread(target=run_process, args=("C", [c_exe], results["C"]))
]

print("🚀 Lancement parallèle des résolutions Go, Python, JavaScript, C...")
for t in threads:
    t.start()
for t in threads:
    t.join()

# Résultats
print("\n📊 Résultats :")
for label, data in results.items():
    if data["time"] < 0:
        print(f"❌ {label} a échoué.")
    else:
        print(f"⏱️ {label:<24}: {data['time']:.4f} s")

print("\n🔎 Sorties :")
for label, data in results.items():
    print(f"--- {label} ---")
    print(data["output"])
    print()

valid = {k: v["time"] for k, v in results.items() if v["time"] >= 0}
if valid:
    winner = min(valid, key=valid.get)
    print(f"🏆 Le plus rapide : {winner}")
else:
    print("❌ Aucun programme n'a terminé correctement.")
