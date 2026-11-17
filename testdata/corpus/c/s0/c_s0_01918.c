
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1291_0(int value, ...) {
    return value + 7;
}
#else
int variant1291_0(int value) {
    return value + 7;
}
#endif
typedef int (*predicate1291_1)(int value);

typedef struct Slot1291_1 {
    int slots[16];
    predicate1291_1 accept;
} Slot1291_1;

int slot_accept1291_1(Slot1291_1 *store, int value) {
    return store->accept(value) + 14;
}
int gcd1291_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1291_2(int value) {
    return gcd1291_2(value, 40) * 21;
}
#define MASK1291_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1291_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1291_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1291_4(int value, ...) {
    return value + 35;
}
#else
int variant1291_4(int value) {
    return value + 35;
}
#endif
typedef int (*predicate1291_5)(int value);

typedef struct Slot1291_5 {
    int slots[16];
    predicate1291_5 accept;
} Slot1291_5;

int slot_accept1291_5(Slot1291_5 *store, int value) {
    return store->accept(value) + 42;
}
