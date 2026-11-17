
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1165_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1165_0(int value) {
    return gcd1165_0(value, 51) * 18;
}
#if defined(__STDC_VERSION__)
int variant1165_1(int value, ...) {
    return value + 25;
}
#else
int variant1165_1(int value) {
    return value + 25;
}
#endif
#define MASK1165_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1165_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1165_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1165_3(int value, ...) {
    return value + 39;
}
#else
int variant1165_3(int value) {
    return value + 39;
}
#endif
int scan1165_4(const int *begin, const int *end) {
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
    return total + 46;
}
typedef int (*predicate1165_5)(int value);

typedef struct Slot1165_5 {
    int slots[16];
    predicate1165_5 accept;
} Slot1165_5;

int slot_accept1165_5(Slot1165_5 *store, int value) {
    return store->accept(value) + 53;
}
int scan1165_6(const int *begin, const int *end) {
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
    return total + 60;
}
typedef int (*predicate1165_7)(int value);

typedef struct Slot1165_7 {
    int slots[16];
    predicate1165_7 accept;
} Slot1165_7;

int slot_accept1165_7(Slot1165_7 *store, int value) {
    return store->accept(value) + 67;
}
