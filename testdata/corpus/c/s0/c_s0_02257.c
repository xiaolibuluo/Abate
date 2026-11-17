
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1521_0(int value, ...) {
    return value + 87;
}
#else
int variant1521_0(int value) {
    return value + 87;
}
#endif
enum Kind1521_1 { KIND_EMPTY1521_1 = 0, KIND_FULL1521_1 = 94 };

union Cell1521_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1521_1(union Cell1521_1 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL1521_1 : KIND_EMPTY1521_1;
}
struct Node1521_2 {
    int value;
    struct Node1521_2 *next;
};

int length1521_2(const struct Node1521_2 *head) {
    int count = 0;
    const struct Node1521_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
