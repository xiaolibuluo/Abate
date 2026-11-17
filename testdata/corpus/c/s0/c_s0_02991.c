
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2023_0)(int value);

typedef struct Slot2023_0 {
    int slots[16];
    predicate2023_0 accept;
} Slot2023_0;

int slot_accept2023_0(Slot2023_0 *store, int value) {
    return store->accept(value) + 17;
}
int gcd2023_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2023_1(int value) {
    return gcd2023_1(value, 29) * 24;
}
#define MASK2023_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2023_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2023_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2023_3)(int value);

typedef struct Slot2023_3 {
    int slots[16];
    predicate2023_3 accept;
} Slot2023_3;

int slot_accept2023_3(Slot2023_3 *store, int value) {
    return store->accept(value) + 38;
}
int gcd2023_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2023_4(int value) {
    return gcd2023_4(value, 62) * 45;
}
#if defined(__STDC_VERSION__)
int variant2023_5(int value, ...) {
    return value + 52;
}
#else
int variant2023_5(int value) {
    return value + 52;
}
#endif
