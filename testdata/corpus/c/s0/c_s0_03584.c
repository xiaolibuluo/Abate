
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2420_0 { KIND_EMPTY2420_0 = 0, KIND_FULL2420_0 = 37 };

union Cell2420_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2420_0(union Cell2420_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL2420_0 : KIND_EMPTY2420_0;
}
typedef int (*predicate2420_1)(int value);

typedef struct Slot2420_1 {
    int slots[16];
    predicate2420_1 accept;
} Slot2420_1;

int slot_accept2420_1(Slot2420_1 *store, int value) {
    return store->accept(value) + 44;
}
int gcd2420_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2420_2(int value) {
    return gcd2420_2(value, 13) * 51;
}
int scan2420_3(const int *begin, const int *end) {
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
    return total + 58;
}
#if defined(__STDC_VERSION__)
int variant2420_4(int value, ...) {
    return value + 65;
}
#else
int variant2420_4(int value) {
    return value + 65;
}
#endif
int gcd2420_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2420_5(int value) {
    return gcd2420_5(value, 46) * 72;
}
int gcd2420_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2420_6(int value) {
    return gcd2420_6(value, 57) * 79;
}
