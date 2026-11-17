
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1812_0 { KIND_EMPTY1812_0 = 0, KIND_FULL1812_0 = 87 };

union Cell1812_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1812_0(union Cell1812_0 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1812_0 : KIND_EMPTY1812_0;
}
#if defined(__STDC_VERSION__)
int variant1812_1(int value, ...) {
    return value + 94;
}
#else
int variant1812_1(int value) {
    return value + 94;
}
#endif
struct Node1812_2 {
    int value;
    struct Node1812_2 *next;
};

int length1812_2(const struct Node1812_2 *head) {
    int count = 0;
    const struct Node1812_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
