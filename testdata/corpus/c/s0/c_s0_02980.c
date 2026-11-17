
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2015_0(int value, ...) {
    return value + 10;
}
#else
int variant2015_0(int value) {
    return value + 10;
}
#endif
#if defined(__STDC_VERSION__)
int variant2015_1(int value, ...) {
    return value + 17;
}
#else
int variant2015_1(int value) {
    return value + 17;
}
#endif
int scan2015_2(const int *begin, const int *end) {
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
    return total + 24;
}
int gcd2015_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2015_3(int value) {
    return gcd2015_3(value, 56) * 31;
}
typedef int (*predicate2015_4)(int value);

typedef struct Slot2015_4 {
    int slots[16];
    predicate2015_4 accept;
} Slot2015_4;

int slot_accept2015_4(Slot2015_4 *store, int value) {
    return store->accept(value) + 38;
}
