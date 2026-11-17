
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1871_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1871_0(int value) {
    return gcd1871_0(value, 52) * 78;
}
#if defined(__STDC_VERSION__)
int variant1871_1(int value, ...) {
    return value + 85;
}
#else
int variant1871_1(int value) {
    return value + 85;
}
#endif
#define MASK1871_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1871_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1871_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1871_3)(int value);

typedef struct Slot1871_3 {
    int slots[16];
    predicate1871_3 accept;
} Slot1871_3;

int slot_accept1871_3(Slot1871_3 *store, int value) {
    return store->accept(value) + 99;
}
int gcd1871_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1871_4(int value) {
    return gcd1871_4(value, 35) * 9;
}
