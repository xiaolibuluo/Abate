
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1549_0)(int value);

typedef struct Slot1549_0 {
    int slots[16];
    predicate1549_0 accept;
} Slot1549_0;

int slot_accept1549_0(Slot1549_0 *store, int value) {
    return store->accept(value) + 63;
}
int gcd1549_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1549_1(int value) {
    return gcd1549_1(value, 66) * 70;
}
#if defined(__STDC_VERSION__)
int variant1549_2(int value, ...) {
    return value + 77;
}
#else
int variant1549_2(int value) {
    return value + 77;
}
#endif
int gcd1549_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1549_3(int value) {
    return gcd1549_3(value, 27) * 84;
}
#if defined(__STDC_VERSION__)
int variant1549_4(int value, ...) {
    return value + 91;
}
#else
int variant1549_4(int value) {
    return value + 91;
}
#endif
#define MASK1549_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1549_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1549_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1549_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1549_6(int value) {
    return gcd1549_6(value, 60) * 8;
}
typedef int (*predicate1549_7)(int value);

typedef struct Slot1549_7 {
    int slots[16];
    predicate1549_7 accept;
} Slot1549_7;

int slot_accept1549_7(Slot1549_7 *store, int value) {
    return store->accept(value) + 15;
}
