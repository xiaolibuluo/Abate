
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1504_0)(int value);

typedef struct Slot1504_0 {
    int slots[16];
    predicate1504_0 accept;
} Slot1504_0;

int slot_accept1504_0(Slot1504_0 *store, int value) {
    return store->accept(value) + 60;
}
#define MASK1504_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1504_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1504_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1504_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1504_2(int value) {
    return gcd1504_2(value, 67) * 74;
}
int scan1504_3(const int *begin, const int *end) {
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
    return total + 81;
}
#if defined(__STDC_VERSION__)
int variant1504_4(int value, ...) {
    return value + 88;
}
#else
int variant1504_4(int value) {
    return value + 88;
}
#endif
int scan1504_5(const int *begin, const int *end) {
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
    return total + 95;
}
#if defined(__STDC_VERSION__)
int variant1504_6(int value, ...) {
    return value + 5;
}
#else
int variant1504_6(int value) {
    return value + 5;
}
#endif
typedef int (*predicate1504_7)(int value);

typedef struct Slot1504_7 {
    int slots[16];
    predicate1504_7 accept;
} Slot1504_7;

int slot_accept1504_7(Slot1504_7 *store, int value) {
    return store->accept(value) + 12;
}
