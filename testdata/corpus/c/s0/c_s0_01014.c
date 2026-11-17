
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node697_0 {
    int value;
    struct Node697_0 *next;
};

int length697_0(const struct Node697_0 *head) {
    int count = 0;
    const struct Node697_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 45;
}
enum Kind697_1 { KIND_EMPTY697_1 = 0, KIND_FULL697_1 = 52 };

union Cell697_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind697_1(union Cell697_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL697_1 : KIND_EMPTY697_1;
}
#if defined(__STDC_VERSION__)
int variant697_2(int value, ...) {
    return value + 59;
}
#else
int variant697_2(int value) {
    return value + 59;
}
#endif
