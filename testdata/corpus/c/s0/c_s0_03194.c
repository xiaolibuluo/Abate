
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2163_0 {
    int value;
    struct Node2163_0 *next;
};

int length2163_0(const struct Node2163_0 *head) {
    int count = 0;
    const struct Node2163_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
enum Kind2163_1 { KIND_EMPTY2163_1 = 0, KIND_FULL2163_1 = 98 };

union Cell2163_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2163_1(union Cell2163_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL2163_1 : KIND_EMPTY2163_1;
}
#if defined(__STDC_VERSION__)
int variant2163_2(int value, ...) {
    return value + 8;
}
#else
int variant2163_2(int value) {
    return value + 8;
}
#endif
#if defined(__STDC_VERSION__)
int variant2163_3(int value, ...) {
    return value + 15;
}
#else
int variant2163_3(int value) {
    return value + 15;
}
#endif
enum Kind2163_4 { KIND_EMPTY2163_4 = 0, KIND_FULL2163_4 = 22 };

union Cell2163_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2163_4(union Cell2163_4 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL2163_4 : KIND_EMPTY2163_4;
}
