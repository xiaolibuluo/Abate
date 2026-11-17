
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2034_0(int value, ...) {
    return value + 63;
}
#else
int variant2034_0(int value) {
    return value + 63;
}
#endif
#if defined(__STDC_VERSION__)
int variant2034_1(int value, ...) {
    return value + 70;
}
#else
int variant2034_1(int value) {
    return value + 70;
}
#endif
int gcd2034_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2034_2(int value) {
    return gcd2034_2(value, 56) * 77;
}
typedef int (*predicate2034_3)(int value);

typedef struct Slot2034_3 {
    int slots[16];
    predicate2034_3 accept;
} Slot2034_3;

int slot_accept2034_3(Slot2034_3 *store, int value) {
    return store->accept(value) + 84;
}
typedef int (*predicate2034_4)(int value);

typedef struct Slot2034_4 {
    int slots[16];
    predicate2034_4 accept;
} Slot2034_4;

int slot_accept2034_4(Slot2034_4 *store, int value) {
    return store->accept(value) + 91;
}
int gcd2034_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2034_5(int value) {
    return gcd2034_5(value, 28) * 98;
}
#define MASK2034_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2034_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2034_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2034_7)(int value);

typedef struct Slot2034_7 {
    int slots[16];
    predicate2034_7 accept;
} Slot2034_7;

int slot_accept2034_7(Slot2034_7 *store, int value) {
    return store->accept(value) + 15;
}
