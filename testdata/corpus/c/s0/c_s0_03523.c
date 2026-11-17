
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2377_0 { KIND_EMPTY2377_0 = 0, KIND_FULL2377_0 = 60 };

union Cell2377_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2377_0(union Cell2377_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL2377_0 : KIND_EMPTY2377_0;
}
int gcd2377_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2377_1(int value) {
    return gcd2377_1(value, 67) * 67;
}
#if defined(__STDC_VERSION__)
int variant2377_2(int value, ...) {
    return value + 74;
}
#else
int variant2377_2(int value) {
    return value + 74;
}
#endif
int gcd2377_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2377_3(int value) {
    return gcd2377_3(value, 28) * 81;
}
#if defined(__STDC_VERSION__)
int variant2377_4(int value, ...) {
    return value + 88;
}
#else
int variant2377_4(int value) {
    return value + 88;
}
#endif
typedef int (*predicate2377_5)(int value);

typedef struct Slot2377_5 {
    int slots[16];
    predicate2377_5 accept;
} Slot2377_5;

int slot_accept2377_5(Slot2377_5 *store, int value) {
    return store->accept(value) + 95;
}
int scan2377_6(const int *begin, const int *end) {
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
    return total + 5;
}
int scan2377_7(const int *begin, const int *end) {
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
    return total + 12;
}
