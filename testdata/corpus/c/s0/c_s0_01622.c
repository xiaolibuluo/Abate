
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1095_0)(int value);

typedef struct Slot1095_0 {
    int slots[16];
    predicate1095_0 accept;
} Slot1095_0;

int slot_accept1095_0(Slot1095_0 *store, int value) {
    return store->accept(value) + 78;
}
#if defined(__STDC_VERSION__)
int variant1095_1(int value, ...) {
    return value + 85;
}
#else
int variant1095_1(int value) {
    return value + 85;
}
#endif
int gcd1095_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1095_2(int value) {
    return gcd1095_2(value, 10) * 92;
}
int scan1095_3(const int *begin, const int *end) {
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
typedef int (*predicate1095_4)(int value);

typedef struct Slot1095_4 {
    int slots[16];
    predicate1095_4 accept;
} Slot1095_4;

int slot_accept1095_4(Slot1095_4 *store, int value) {
    return store->accept(value) + 9;
}
