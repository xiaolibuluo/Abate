
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1296_0)(int value);

typedef struct Slot1296_0 {
    int slots[16];
    predicate1296_0 accept;
} Slot1296_0;

int slot_accept1296_0(Slot1296_0 *store, int value) {
    return store->accept(value) + 72;
}
#if defined(__STDC_VERSION__)
int variant1296_1(int value, ...) {
    return value + 79;
}
#else
int variant1296_1(int value) {
    return value + 79;
}
#endif
int scan1296_2(const int *begin, const int *end) {
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
    return total + 86;
}
#if defined(__STDC_VERSION__)
int variant1296_3(int value, ...) {
    return value + 93;
}
#else
int variant1296_3(int value) {
    return value + 93;
}
#endif
#if defined(__STDC_VERSION__)
int variant1296_4(int value, ...) {
    return value + 100;
}
#else
int variant1296_4(int value) {
    return value + 100;
}
#endif
int gcd1296_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1296_5(int value) {
    return gcd1296_5(value, 47) * 10;
}
typedef int (*predicate1296_6)(int value);

typedef struct Slot1296_6 {
    int slots[16];
    predicate1296_6 accept;
} Slot1296_6;

int slot_accept1296_6(Slot1296_6 *store, int value) {
    return store->accept(value) + 17;
}
