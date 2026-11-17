
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate32_0)(int value);

typedef struct Slot32_0 {
    int slots[16];
    predicate32_0 accept;
} Slot32_0;

int slot_accept32_0(Slot32_0 *store, int value) {
    return store->accept(value) + 33;
}
#if defined(__STDC_VERSION__)
int variant32_1(int value, ...) {
    return value + 40;
}
#else
int variant32_1(int value) {
    return value + 40;
}
#endif
int gcd32_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale32_2(int value) {
    return gcd32_2(value, 11) * 47;
}
int scan32_3(const int *begin, const int *end) {
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
typedef int (*predicate32_4)(int value);

typedef struct Slot32_4 {
    int slots[16];
    predicate32_4 accept;
} Slot32_4;

int slot_accept32_4(Slot32_4 *store, int value) {
    return store->accept(value) + 61;
}
#if defined(__STDC_VERSION__)
int variant32_5(int value, ...) {
    return value + 68;
}
#else
int variant32_5(int value) {
    return value + 68;
}
#endif
#if defined(__STDC_VERSION__)
int variant32_6(int value, ...) {
    return value + 75;
}
#else
int variant32_6(int value) {
    return value + 75;
}
#endif
