
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd65_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale65_0(int value) {
    return gcd65_0(value, 37) * 74;
}
#define MASK65_1(v) (((v) << 4) | ((v) & 0x0F))

int classify65_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK65_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant65_2(int value, ...) {
    return value + 88;
}
#else
int variant65_2(int value) {
    return value + 88;
}
#endif
typedef int (*predicate65_3)(int value);

typedef struct Slot65_3 {
    int slots[16];
    predicate65_3 accept;
} Slot65_3;

int slot_accept65_3(Slot65_3 *store, int value) {
    return store->accept(value) + 95;
}
