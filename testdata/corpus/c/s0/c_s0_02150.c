
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1447_0(int value, ...) {
    return value + 95;
}
#else
int variant1447_0(int value) {
    return value + 95;
}
#endif
int scan1447_1(const int *begin, const int *end) {
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
int gcd1447_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1447_2(int value) {
    return gcd1447_2(value, 34) * 12;
}
typedef int (*predicate1447_3)(int value);

typedef struct Slot1447_3 {
    int slots[16];
    predicate1447_3 accept;
} Slot1447_3;

int slot_accept1447_3(Slot1447_3 *store, int value) {
    return store->accept(value) + 19;
}
#if defined(__STDC_VERSION__)
int variant1447_4(int value, ...) {
    return value + 26;
}
#else
int variant1447_4(int value) {
    return value + 26;
}
#endif
enum Kind1447_5 { KIND_EMPTY1447_5 = 0, KIND_FULL1447_5 = 33 };

union Cell1447_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1447_5(union Cell1447_5 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL1447_5 : KIND_EMPTY1447_5;
}
int scan1447_6(const int *begin, const int *end) {
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
    return total + 40;
}
