
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1402_0)(int value);

typedef struct Slot1402_0 {
    int slots[16];
    predicate1402_0 accept;
} Slot1402_0;

int slot_accept1402_0(Slot1402_0 *store, int value) {
    return store->accept(value) + 92;
}
#if defined(__STDC_VERSION__)
int variant1402_1(int value, ...) {
    return value + 99;
}
#else
int variant1402_1(int value) {
    return value + 99;
}
#endif
int scan1402_2(const int *begin, const int *end) {
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
int scan1402_3(const int *begin, const int *end) {
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
    return total + 16;
}
typedef int (*predicate1402_4)(int value);

typedef struct Slot1402_4 {
    int slots[16];
    predicate1402_4 accept;
} Slot1402_4;

int slot_accept1402_4(Slot1402_4 *store, int value) {
    return store->accept(value) + 23;
}
int gcd1402_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1402_5(int value) {
    return gcd1402_5(value, 57) * 30;
}
#define MASK1402_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1402_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1402_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
