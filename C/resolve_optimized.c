#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "cJSON.h"
#include <cblas.h>

#define SIZE 1000

double *flatten_matrix(double **matrix) {
    double *flat = malloc(SIZE * SIZE * sizeof(double));
    for (int i = 0; i < SIZE; i++)
        for (int j = 0; j < SIZE; j++)
            flat[i * SIZE + j] = matrix[i][j];
    return flat;
}

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
    double *flatA = flatten_matrix(A);
    double *flatB = flatten_matrix(B);
    double *result = calloc(SIZE * SIZE, sizeof(double));

    printf("[INFO] Multiplication avec OpenBLAS...\n");
    clock_t start = clock();
    cblas_dgemm(CblasRowMajor, CblasNoTrans, CblasNoTrans,
                SIZE, SIZE, SIZE, 1.0,
                flatA, SIZE,
                flatB, SIZE,
                0.0, result, SIZE);
    clock_t end = clock();

    double elapsed = (double)(end - start) / CLOCKS_PER_SEC;
    printf("[RESULT] Temps de résolution (C + OpenBLAS) : %.4f secondes\n", elapsed);

    // Clean
    free(data);
    cJSON_Delete(json);
    free_matrix(A);
    free_matrix(B);
    free(flatA);
    free(flatB);
    free(result);

    return 0;
}
