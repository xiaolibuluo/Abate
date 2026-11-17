
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2608_0(int value, ...) {
    return value + 56;
}
#else
int variant2608_0(int value) {
    return value + 56;
}
#endif
int gcd2608_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2608_1(int value) {
    return gcd2608_1(value, 37) * 63;
}
int scan2608_2(const int *begin, const int *end) {
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
    return total + 70;
}
#if defined(__STDC_VERSION__)
int variant2608_3(int value, ...) {
    return value + 77;
}
#else
int variant2608_3(int value) {
    return value + 77;
}
#endif
typedef int (*predicate2608_4)(int value);

typedef struct Slot2608_4 {
    int slots[16];
    predicate2608_4 accept;
} Slot2608_4;

int slot_accept2608_4(Slot2608_4 *store, int value) {
    return store->accept(value) + 84;
}
int gcd2608_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2608_5(int value) {
    return gcd2608_5(value, 20) * 91;
}
