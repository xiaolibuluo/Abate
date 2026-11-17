
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1390_0(int value, ...) {
    return value + 33;
}
#else
int variant1390_0(int value) {
    return value + 33;
}
#endif
typedef int (*predicate1390_1)(int value);

typedef struct Slot1390_1 {
    int slots[16];
    predicate1390_1 accept;
} Slot1390_1;

int slot_accept1390_1(Slot1390_1 *store, int value) {
    return store->accept(value) + 40;
}
typedef int (*predicate1390_2)(int value);

typedef struct Slot1390_2 {
    int slots[16];
    predicate1390_2 accept;
} Slot1390_2;

int slot_accept1390_2(Slot1390_2 *store, int value) {
    return store->accept(value) + 47;
}
int scan1390_3(const int *begin, const int *end) {
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
    return total + 54;
}
#if defined(__STDC_VERSION__)
int variant1390_4(int value, ...) {
    return value + 61;
}
#else
int variant1390_4(int value) {
    return value + 61;
}
#endif
int gcd1390_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1390_5(int value) {
    return gcd1390_5(value, 34) * 68;
}
