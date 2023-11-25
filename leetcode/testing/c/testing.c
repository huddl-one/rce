#include <stdio.h>
#include <stdlib.h>

void daily_temperatures(int* temperatures, int n, int* result);

typedef struct {
    int* input;
    int inputSize;
    int* expectedResult;
    int expectedSize;
} TestCase;

void printArray(int *arr, int size) {
    printf("[");
    for (int i = 0; i < size; ++i) {
        printf("%d", arr[i]);
        if (i < size - 1) printf(", ");
    }
    printf("]");
}

void runTest(TestCase test, void (*function)(int*, int, int*)) {
    int* result = malloc(test.expectedSize * sizeof(int));
    function(test.input, test.inputSize, result);

    printf("{\n  \"input\": ");
    printArray(test.input, test.inputSize);
    printf(",\n  \"expected\": ");
    printArray(test.expectedResult, test.expectedSize);
    printf(",\n  \"output\": ");
    printArray(result, test.expectedSize);

    int passed = 1;
    for (int i = 0; i < test.expectedSize; i++) {
        if (result[i] != test.expectedResult[i]) {
            passed = 0;
            break;
        }
    }
    printf(",\n  \"passed\": %s\n}\n", passed ? "true" : "false");

    free(result);
}

int main() {
    // Define an array of test cases
    TestCase tests[] = {
        { (int[]){73, 74, 75, 71, 69, 72, 76, 73}, 8, (int[]){1, 1, 4, 2, 1, 1, 0, 0}, 8 },
        { (int[]){30, 40, 50, 60}, 4, (int[]){1, 1, 1, 0}, 4 },
        // Add more test cases here
    };
    int numTests = sizeof(tests) / sizeof(tests[0]);

    printf("[\n");
    for (int i = 0; i < numTests; ++i) {
        runTest(tests[i], daily_temperatures);
        if (i < numTests - 1) printf(",\n");
    }
    printf("\n]\n");

    return 0;
}
