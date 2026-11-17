
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2568_0(int value, ...) {
    return value + 21;
}
#else
int variant2568_0(int value) {
    return value + 21;
}
#endif
int gcd2568_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2568_1(int value) {
    return gcd2568_1(value, 62) * 28;
}
#if defined(__STDC_VERSION__)
int variant2568_2(int value, ...) {
    return value + 35;
}
#else
int variant2568_2(int value) {
    return value + 35;
}
#endif
#if defined(__STDC_VERSION__)
int variant2568_3(int value, ...) {
    return value + 42;
}
#else
int variant2568_3(int value) {
    return value + 42;
}
#endif
int scan2568_4(const int *begin, const int *end) {
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
    return total + 49;
}
#define MASK2568_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2568_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2568_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2568_6)(int value);

typedef struct Slot2568_6 {
    int slots[16];
    predicate2568_6 accept;
} Slot2568_6;

int slot_accept2568_6(Slot2568_6 *store, int value) {
    return store->accept(value) + 63;
}
int gcd2568_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2568_7(int value) {
    return gcd2568_7(value, 67) * 70;
}
