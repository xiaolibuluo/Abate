
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1986_0(int value, ...) {
    return value + 21;
}
#else
int variant1986_0(int value) {
    return value + 21;
}
#endif
#define MASK1986_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1986_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1986_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1986_2)(int value);

typedef struct Slot1986_2 {
    int slots[16];
    predicate1986_2 accept;
} Slot1986_2;

int slot_accept1986_2(Slot1986_2 *store, int value) {
    return store->accept(value) + 35;
}
int gcd1986_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1986_3(int value) {
    return gcd1986_3(value, 36) * 42;
}
#if defined(__STDC_VERSION__)
int variant1986_4(int value, ...) {
    return value + 49;
}
#else
int variant1986_4(int value) {
    return value + 49;
}
#endif
int scan1986_5(const int *begin, const int *end) {
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
    return total + 56;
}
int gcd1986_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1986_6(int value) {
    return gcd1986_6(value, 69) * 63;
}
#if defined(__STDC_VERSION__)
int variant1986_7(int value, ...) {
    return value + 70;
}
#else
int variant1986_7(int value) {
    return value + 70;
}
#endif
