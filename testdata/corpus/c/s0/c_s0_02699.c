
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1835_0(int value, ...) {
    return value + 95;
}
#else
int variant1835_0(int value) {
    return value + 95;
}
#endif
typedef int (*predicate1835_1)(int value);

typedef struct Slot1835_1 {
    int slots[16];
    predicate1835_1 accept;
} Slot1835_1;

int slot_accept1835_1(Slot1835_1 *store, int value) {
    return store->accept(value) + 5;
}
#if defined(__STDC_VERSION__)
int variant1835_2(int value, ...) {
    return value + 12;
}
#else
int variant1835_2(int value) {
    return value + 12;
}
#endif
#if defined(__STDC_VERSION__)
int variant1835_3(int value, ...) {
    return value + 19;
}
#else
int variant1835_3(int value) {
    return value + 19;
}
#endif
int gcd1835_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1835_4(int value) {
    return gcd1835_4(value, 27) * 26;
}
#define MASK1835_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1835_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1835_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1835_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1835_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1835_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1835_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1835_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1835_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
