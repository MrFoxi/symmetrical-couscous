#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "cJSON.h"

#define SIZE 1000

double **allocate_matrix() {
    double **matrix = malloc(SIZE * sizeof(double *));
    for (int i = 0; i < SIZE; i++) {
        matrix[i] = malloc(SIZE * sizeof(double));
    }
    return matrix;
}

void free_matrix(double **matrix) {
    for (int i = 0; i < SIZE; i++) {
        free(matrix[i]);
    }
    free(matrix);
}

void multiply_matrices(double **a, double **b, double **result) {
    for (int i = 0; i < SIZE; i++) {
        for (int j = 0; j < SIZE; j++) {
            double sum = 0.0;
            for (int k = 0; k < SIZE; k++) {
                sum += a[i][k] * b[k][j];
            }
            result[i][j] = sum;
        }
    }
}

double **parse_matrix(cJSON *json_matrix) {
    double **matrix = allocate_matrix();
    int i = 0;
    cJSON *row;
    cJSON_ArrayForEach(row, json_matrix) {
        int j = 0;
        cJSON *val;
        cJSON_ArrayForEach(val, row) {
            matrix[i][j++] = val->valuedouble;
        }
        i++;
    }
    return matrix;
}

int main() {
    printf("[INFO] Chargement des matrices depuis matrices.json...\n");

    FILE *f = fopen("matrices.json", "rb");
    if (!f) {
        printf("Erreur ouverture du fichier.\n");
        return 1;
    }

    fseek(f, 0, SEEK_END);
    long len = ftell(f);
    fseek(f, 0, SEEK_SET);

    char *data = malloc(len + 1);
    fread(data, 1, len, f);
    data[len] = '\0';
    fclose(f);

    cJSON *json = cJSON_Parse(data);
    if (!json) {
        printf("Erreur JSON.\n");
        return 1;
    }

    double **A = parse_matrix(cJSON_GetObjectItem(json, "A"));
    double **B = parse_matrix(cJSON_GetObjectItem(json, "B"));
    double **result = allocate_matrix();

    printf("[INFO] Multiplication (C)...\n");
    clock_t start = clock();
    multiply_matrices(A, B, result);
    clock_t end = clock();

    double elapsed = (double)(end - start) / CLOCKS_PER_SEC;
    printf("[RESULT] Temps de résolution (C) : %.4f secondes\n", elapsed);

    free(data);
    cJSON_Delete(json);
    free_matrix(A);
    free_matrix(B);
    free_matrix(result);

    return 0;
}
