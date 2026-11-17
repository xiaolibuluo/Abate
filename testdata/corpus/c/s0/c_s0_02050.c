
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1386_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1386_0(int value) {
    return gcd1386_0(value, 12) * 78;
}
#if defined(__STDC_VERSION__)
int variant1386_1(int value, ...) {
    return value + 85;
}
#else
int variant1386_1(int value) {
    return value + 85;
}
#endif
typedef int (*predicate1386_2)(int value);

typedef struct Slot1386_2 {
    int slots[16];
    predicate1386_2 accept;
} Slot1386_2;

int slot_accept1386_2(Slot1386_2 *store, int value) {
    return store->accept(value) + 92;
}
int scan1386_3(const int *begin, const int *end) {
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
    return total + 99;
}
int scan1386_4(const int *begin, const int *end) {
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
    return total + 9;
}
