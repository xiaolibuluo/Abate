
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1437_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1437_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1437_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1437_1)(int value);

typedef struct Slot1437_1 {
    int slots[16];
    predicate1437_1 accept;
} Slot1437_1;

int slot_accept1437_1(Slot1437_1 *store, int value) {
    return store->accept(value) + 69;
}
typedef int (*predicate1437_2)(int value);

typedef struct Slot1437_2 {
    int slots[16];
    predicate1437_2 accept;
} Slot1437_2;

int slot_accept1437_2(Slot1437_2 *store, int value) {
    return store->accept(value) + 76;
}
#if defined(__STDC_VERSION__)
int variant1437_3(int value, ...) {
    return value + 83;
}
#else
int variant1437_3(int value) {
    return value + 83;
}
#endif
typedef int (*predicate1437_4)(int value);

typedef struct Slot1437_4 {
    int slots[16];
    predicate1437_4 accept;
} Slot1437_4;

int slot_accept1437_4(Slot1437_4 *store, int value) {
    return store->accept(value) + 90;
}
int gcd1437_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1437_5(int value) {
    return gcd1437_5(value, 58) * 97;
}
