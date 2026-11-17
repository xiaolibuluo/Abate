
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2553_0(int value, ...) {
    return value + 20;
}
#else
int variant2553_0(int value) {
    return value + 20;
}
#endif
#if defined(__STDC_VERSION__)
int variant2553_1(int value, ...) {
    return value + 27;
}
#else
int variant2553_1(int value) {
    return value + 27;
}
#endif
int gcd2553_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2553_2(int value) {
    return gcd2553_2(value, 29) * 34;
}
typedef int (*predicate2553_3)(int value);

typedef struct Slot2553_3 {
    int slots[16];
    predicate2553_3 accept;
} Slot2553_3;

int slot_accept2553_3(Slot2553_3 *store, int value) {
    return store->accept(value) + 41;
}
#define MASK2553_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2553_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2553_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
