
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1280_0(int value, ...) {
    return value + 58;
}
#else
int variant1280_0(int value) {
    return value + 58;
}
#endif
struct Node1280_1 {
    int value;
    struct Node1280_1 *next;
};

int length1280_1(const struct Node1280_1 *head) {
    int count = 0;
    const struct Node1280_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
enum Kind1280_2 { KIND_EMPTY1280_2 = 0, KIND_FULL1280_2 = 72 };

union Cell1280_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1280_2(union Cell1280_2 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1280_2 : KIND_EMPTY1280_2;
}
