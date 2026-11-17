
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1199_0(int value, ...) {
    return value + 72;
}
#else
int variant1199_0(int value) {
    return value + 72;
}
#endif
#if defined(__STDC_VERSION__)
int variant1199_1(int value, ...) {
    return value + 79;
}
#else
int variant1199_1(int value) {
    return value + 79;
}
#endif
int gcd1199_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1199_2(int value) {
    return gcd1199_2(value, 67) * 86;
}
#if defined(__STDC_VERSION__)
int variant1199_3(int value, ...) {
    return value + 93;
}
#else
int variant1199_3(int value) {
    return value + 93;
}
#endif
#if defined(__STDC_VERSION__)
int variant1199_4(int value, ...) {
    return value + 100;
}
#else
int variant1199_4(int value) {
    return value + 100;
}
#endif
#define MASK1199_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1199_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1199_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1199_6)(int value);

typedef struct Slot1199_6 {
    int slots[16];
    predicate1199_6 accept;
} Slot1199_6;

int slot_accept1199_6(Slot1199_6 *store, int value) {
    return store->accept(value) + 17;
}
