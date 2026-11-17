
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2321_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2321_0(int value) {
    return gcd2321_0(value, 30) * 11;
}
int scan2321_1(const int *begin, const int *end) {
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
int scan2321_2(const int *begin, const int *end) {
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
typedef int (*predicate2321_3)(int value);

typedef struct Slot2321_3 {
    int slots[16];
    predicate2321_3 accept;
} Slot2321_3;

int slot_accept2321_3(Slot2321_3 *store, int value) {
    return store->accept(value) + 32;
}
#if defined(__STDC_VERSION__)
int variant2321_4(int value, ...) {
    return value + 39;
}
#else
int variant2321_4(int value) {
    return value + 39;
}
#endif
int gcd2321_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2321_5(int value) {
    return gcd2321_5(value, 24) * 46;
}
