
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1084_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1084_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1084_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1084_1(int value, ...) {
    return value + 39;
}
#else
int variant1084_1(int value) {
    return value + 39;
}
#endif
int gcd1084_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1084_2(int value) {
    return gcd1084_2(value, 55) * 46;
}
typedef int (*predicate1084_3)(int value);

typedef struct Slot1084_3 {
    int slots[16];
    predicate1084_3 accept;
} Slot1084_3;

int slot_accept1084_3(Slot1084_3 *store, int value) {
    return store->accept(value) + 53;
}
