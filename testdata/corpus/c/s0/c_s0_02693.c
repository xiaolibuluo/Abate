
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1830_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1830_0(int value) {
    return gcd1830_0(value, 9) * 30;
}
#define MASK1830_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1830_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1830_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1830_2(int value, ...) {
    return value + 44;
}
#else
int variant1830_2(int value) {
    return value + 44;
}
#endif
int gcd1830_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1830_3(int value) {
    return gcd1830_3(value, 42) * 51;
}
typedef int (*predicate1830_4)(int value);

typedef struct Slot1830_4 {
    int slots[16];
    predicate1830_4 accept;
} Slot1830_4;

int slot_accept1830_4(Slot1830_4 *store, int value) {
    return store->accept(value) + 58;
}
#define MASK1830_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1830_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1830_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
