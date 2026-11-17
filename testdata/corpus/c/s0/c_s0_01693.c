
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1137_0 { KIND_EMPTY1137_0 = 0, KIND_FULL1137_0 = 42 };

union Cell1137_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1137_0(union Cell1137_0 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1137_0 : KIND_EMPTY1137_0;
}
#if defined(__STDC_VERSION__)
int variant1137_1(int value, ...) {
    return value + 49;
}
#else
int variant1137_1(int value) {
    return value + 49;
}
#endif
struct Node1137_2 {
    int value;
    struct Node1137_2 *next;
};

int length1137_2(const struct Node1137_2 *head) {
    int count = 0;
    const struct Node1137_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
struct Node1137_3 {
    int value;
    struct Node1137_3 *next;
};

int length1137_3(const struct Node1137_3 *head) {
    int count = 0;
    const struct Node1137_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
