#ifndef cJSON__h
#define cJSON__h

#ifdef __cplusplus
extern "C" {
#endif

#include <stddef.h>

typedef struct cJSON {
    struct cJSON *next, *prev;
    struct cJSON *child;

    int type;
    char *valuestring;
    int valueint;
    double valuedouble;

    char *string;
} cJSON;

cJSON *cJSON_Parse(const char *value);
void   cJSON_Delete(cJSON *c);
cJSON *cJSON_GetObjectItem(const cJSON *object, const char *string);
int    cJSON_GetArraySize(const cJSON *array);
cJSON *cJSON_GetArrayItem(const cJSON *array, int index);

#define cJSON_Array 1

#define cJSON_ArrayForEach(element, array)     for (int __i = 0; __i < cJSON_GetArraySize(array) && (element = cJSON_GetArrayItem(array, __i)); __i++)

#ifdef __cplusplus
}
#endif

#endif
