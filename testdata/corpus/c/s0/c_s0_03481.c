
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2350_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2350_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2350_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2350_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2350_1(int value) {
    return gcd2350_1(value, 61) * 7;
}
#if defined(__STDC_VERSION__)
int variant2350_2(int value, ...) {
    return value + 14;
}
#else
int variant2350_2(int value) {
    return value + 14;
}
#endif
typedef int (*predicate2350_3)(int value);

typedef struct Slot2350_3 {
    int slots[16];
    predicate2350_3 accept;
} Slot2350_3;

int slot_accept2350_3(Slot2350_3 *store, int value) {
    return store->accept(value) + 21;
}
