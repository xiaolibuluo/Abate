
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1590_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1590_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1590_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1590_1(int value, ...) {
    return value + 21;
}
#else
int variant1590_1(int value) {
    return value + 21;
}
#endif
int gcd1590_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1590_2(int value) {
    return gcd1590_2(value, 59) * 28;
}
typedef int (*predicate1590_3)(int value);

typedef struct Slot1590_3 {
    int slots[16];
    predicate1590_3 accept;
} Slot1590_3;

int slot_accept1590_3(Slot1590_3 *store, int value) {
    return store->accept(value) + 35;
}
#define MASK1590_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1590_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1590_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1590_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1590_5(int value) {
    return gcd1590_5(value, 31) * 49;
}
typedef int (*predicate1590_6)(int value);

typedef struct Slot1590_6 {
    int slots[16];
    predicate1590_6 accept;
} Slot1590_6;

int slot_accept1590_6(Slot1590_6 *store, int value) {
    return store->accept(value) + 56;
}
