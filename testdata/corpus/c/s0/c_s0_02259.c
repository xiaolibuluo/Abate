
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1523_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1523_0(int value) {
    return gcd1523_0(value, 56) * 16;
}
#if defined(__STDC_VERSION__)
int variant1523_1(int value, ...) {
    return value + 23;
}
#else
int variant1523_1(int value) {
    return value + 23;
}
#endif
#define MASK1523_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1523_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1523_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1523_3)(int value);

typedef struct Slot1523_3 {
    int slots[16];
    predicate1523_3 accept;
} Slot1523_3;

int slot_accept1523_3(Slot1523_3 *store, int value) {
    return store->accept(value) + 37;
}
