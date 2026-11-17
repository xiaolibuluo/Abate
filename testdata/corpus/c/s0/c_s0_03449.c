
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2332_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 57;
}
typedef int (*predicate2332_1)(int value);

typedef struct Slot2332_1 {
    int slots[16];
    predicate2332_1 accept;
} Slot2332_1;

int slot_accept2332_1(Slot2332_1 *store, int value) {
    return store->accept(value) + 64;
}
#if defined(__STDC_VERSION__)
int variant2332_2(int value, ...) {
    return value + 71;
}
#else
int variant2332_2(int value) {
    return value + 71;
}
#endif
enum Kind2332_3 { KIND_EMPTY2332_3 = 0, KIND_FULL2332_3 = 78 };

union Cell2332_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2332_3(union Cell2332_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL2332_3 : KIND_EMPTY2332_3;
}
int gcd2332_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2332_4(int value) {
    return gcd2332_4(value, 29) * 85;
}
#if defined(__STDC_VERSION__)
int variant2332_5(int value, ...) {
    return value + 92;
}
#else
int variant2332_5(int value) {
    return value + 92;
}
#endif
#if defined(__STDC_VERSION__)
int variant2332_6(int value, ...) {
    return value + 99;
}
#else
int variant2332_6(int value) {
    return value + 99;
}
#endif
int gcd2332_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2332_7(int value) {
    return gcd2332_7(value, 62) * 9;
}
#if defined(__STDC_VERSION__)
int variant2332_8(int value, ...) {
    return value + 16;
}
#else
int variant2332_8(int value) {
    return value + 16;
}
#endif
