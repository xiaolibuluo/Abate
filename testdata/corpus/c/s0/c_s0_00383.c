
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind259_0 { KIND_EMPTY259_0 = 0, KIND_FULL259_0 = 74 };

union Cell259_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind259_0(union Cell259_0 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL259_0 : KIND_EMPTY259_0;
}
struct Node259_1 {
    int value;
    struct Node259_1 *next;
};

int length259_1(const struct Node259_1 *head) {
    int count = 0;
    const struct Node259_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
#if defined(__STDC_VERSION__)
int variant259_2(int value, ...) {
    return value + 88;
}
#else
int variant259_2(int value) {
    return value + 88;
}
#endif
