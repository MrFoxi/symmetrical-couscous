#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "cJSON.h"

// Mini implémentation lisant uniquement un tableau 2D de doubles

#include <ctype.h>

static char *skip_whitespace(char *c) {
    while (*c && isspace((unsigned char)*c)) c++;
    return c;
}

cJSON *cJSON_Parse(const char *value) {
    cJSON *root = malloc(sizeof(cJSON));
    memset(root, 0, sizeof(cJSON));
    root->type = cJSON_Array;

    const char *c = value;
    c = strchr(c, '['); // chercher le premier [
    if (!c) return NULL;
    c++;

    cJSON *last_row = NULL;
    while (*c && *c != ']') {
        while (*c && *c != '[') c++; // début de ligne
        if (*c == 0) break;
        c++;

        cJSON *row = malloc(sizeof(cJSON));
        memset(row, 0, sizeof(cJSON));
        row->type = cJSON_Array;

        cJSON *last_val = NULL;
        while (*c && *c != ']') {
            char *end;
            double val = strtod(c, &end);
            if (c == end) break;

            cJSON *item = malloc(sizeof(cJSON));
            memset(item, 0, sizeof(cJSON));
            item->valuedouble = val;
            item->type = 0;

            if (!row->child) {
                row->child = item;
            } else {
                last_val->next = item;
                item->prev = last_val;
            }
            last_val = item;

            c = end;
            while (*c && (*c == ',' || isspace((unsigned char)*c))) c++;
        }

        if (!root->child) {
            root->child = row;
        } else {
            last_row->next = row;
            row->prev = last_row;
        }
        last_row = row;

        while (*c && *c != ']') c++;
        if (*c) c++;
        while (*c && (*c == ',' || isspace((unsigned char)*c))) c++;
    }

    return root;
}

void cJSON_Delete(cJSON *c) {
    if (!c) return;
    if (c->child) {
        cJSON *n = c->child;
        while (n) {
            cJSON *next = n->next;
            cJSON_Delete(n);
            n = next;
        }
    }
    free(c);
}

cJSON *cJSON_GetObjectItem(const cJSON *object, const char *string) {
    if (strcmp(string, "A") == 0) return object->child;
    if (strcmp(string, "B") == 0) return object->child->next;
    return NULL;
}

int cJSON_GetArraySize(const cJSON *array) {
    int count = 0;
    for (cJSON *el = array->child; el; el = el->next) count++;
    return count;
}

cJSON *cJSON_GetArrayItem(const cJSON *array, int index) {
    cJSON *el = array->child;
    while (el && index--) el = el->next;
    return el;
}
