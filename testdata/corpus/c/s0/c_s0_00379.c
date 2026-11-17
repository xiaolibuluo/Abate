
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant255_0(int value, ...) {
    return value + 22;
}
#else
int variant255_0(int value) {
    return value + 22;
}
#endif
enum Kind255_1 { KIND_EMPTY255_1 = 0, KIND_FULL255_1 = 29 };

union Cell255_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind255_1(union Cell255_1 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL255_1 : KIND_EMPTY255_1;
}
struct Node255_2 {
    int value;
    struct Node255_2 *next;
};

int length255_2(const struct Node255_2 *head) {
    int count = 0;
    const struct Node255_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 36;
}
