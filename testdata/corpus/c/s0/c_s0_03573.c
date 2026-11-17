
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2410_0(int value, ...) {
    return value + 101;
}
#else
int variant2410_0(int value) {
    return value + 101;
}
#endif
int gcd2410_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2410_1(int value) {
    return gcd2410_1(value, 54) * 11;
}
int scan2410_2(const int *begin, const int *end) {
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
    return total + 18;
}
int scan2410_3(const int *begin, const int *end) {
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
    return total + 25;
}
typedef int (*predicate2410_4)(int value);

typedef struct Slot2410_4 {
    int slots[16];
    predicate2410_4 accept;
} Slot2410_4;

int slot_accept2410_4(Slot2410_4 *store, int value) {
    return store->accept(value) + 32;
}
