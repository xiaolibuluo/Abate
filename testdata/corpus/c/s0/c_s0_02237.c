
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1505_0)(int value);

typedef struct Slot1505_0 {
    int slots[16];
    predicate1505_0 accept;
} Slot1505_0;

int slot_accept1505_0(Slot1505_0 *store, int value) {
    return store->accept(value) + 73;
}
typedef int (*predicate1505_1)(int value);

typedef struct Slot1505_1 {
    int slots[16];
    predicate1505_1 accept;
} Slot1505_1;

int slot_accept1505_1(Slot1505_1 *store, int value) {
    return store->accept(value) + 80;
}
#if defined(__STDC_VERSION__)
int variant1505_2(int value, ...) {
    return value + 87;
}
#else
int variant1505_2(int value) {
    return value + 87;
}
#endif
int scan1505_3(const int *begin, const int *end) {
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
    return total + 94;
}
int scan1505_4(const int *begin, const int *end) {
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
    return total + 101;
}
int scan1505_5(const int *begin, const int *end) {
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
    return total + 11;
}
int gcd1505_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1505_6(int value) {
    return gcd1505_6(value, 57) * 18;
}
#define MASK1505_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1505_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1505_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
