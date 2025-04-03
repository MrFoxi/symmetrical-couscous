const fs = require('fs');
const math = require('mathjs');

console.log("[INFO] Chargement des matrices depuis matrices.json...");
const raw = fs.readFileSync("matrices.json");
const data = JSON.parse(raw);
const A = data["A"];
const B = data["B"];

console.log("[INFO] Multiplication (JavaScript + mathjs)...");
const start = Date.now();
const result = math.multiply(A, B);
const elapsed = (Date.now() - start) / 1000;

console.log(`[RESULT] Temps de résolution (JavaScript + mathjs) : ${elapsed.toFixed(4)} secondes`);
