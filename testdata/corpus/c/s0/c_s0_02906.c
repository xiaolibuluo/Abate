
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1971_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1971_0(int value) {
    return gcd1971_0(value, 20) * 20;
}
#if defined(__STDC_VERSION__)
int variant1971_1(int value, ...) {
    return value + 27;
}
#else
int variant1971_1(int value) {
    return value + 27;
}
#endif
enum Kind1971_2 { KIND_EMPTY1971_2 = 0, KIND_FULL1971_2 = 34 };

union Cell1971_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1971_2(union Cell1971_2 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1971_2 : KIND_EMPTY1971_2;
}
int scan1971_3(const int *begin, const int *end) {
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
    return total + 41;
}
typedef int (*predicate1971_4)(int value);

typedef struct Slot1971_4 {
    int slots[16];
    predicate1971_4 accept;
} Slot1971_4;

int slot_accept1971_4(Slot1971_4 *store, int value) {
    return store->accept(value) + 48;
}
