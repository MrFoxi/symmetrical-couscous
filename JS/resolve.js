const fs = require('fs');

function multiplyMatrices(a, b) {
    const size = a.length;
    const result = Array.from({ length: size }, () => Array(size).fill(0));

    for (let i = 0; i < size; i++) {
        for (let j = 0; j < size; j++) {
            let sum = 0;
            for (let k = 0; k < size; k++) {
                sum += a[i][k] * b[k][j];
            }
            result[i][j] = sum;
        }
    }
    return result;
}

console.log("[INFO] Chargement des matrices depuis matrices.json...");
const raw = fs.readFileSync("matrices.json");
const data = JSON.parse(raw);
const A = data["A"];
const B = data["B"];

console.log("[INFO] Multiplication (JavaScript pur)...");
const start = Date.now();
multiplyMatrices(A, B);
const elapsed = (Date.now() - start) / 1000;

console.log(`[RESULT] Temps de résolution (JavaScript) : ${elapsed.toFixed(4)} secondes`);
