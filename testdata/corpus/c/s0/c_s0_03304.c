
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2233_0)(int value);

typedef struct Slot2233_0 {
    int slots[16];
    predicate2233_0 accept;
} Slot2233_0;

int slot_accept2233_0(Slot2233_0 *store, int value) {
    return store->accept(value) + 31;
}
#if defined(__STDC_VERSION__)
int variant2233_1(int value, ...) {
    return value + 38;
}
#else
int variant2233_1(int value) {
    return value + 38;
}
#endif
int scan2233_2(const int *begin, const int *end) {
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
    return total + 45;
}
int gcd2233_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2233_3(int value) {
    return gcd2233_3(value, 57) * 52;
}
int scan2233_4(const int *begin, const int *end) {
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
    return total + 59;
}
int scan2233_5(const int *begin, const int *end) {
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
    return total + 66;
}
enum Kind2233_6 { KIND_EMPTY2233_6 = 0, KIND_FULL2233_6 = 73 };

union Cell2233_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2233_6(union Cell2233_6 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL2233_6 : KIND_EMPTY2233_6;
}
