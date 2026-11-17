
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK799_0(v) (((v) << 4) | ((v) & 0x0F))

int classify799_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK799_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate799_1)(int value);

typedef struct Slot799_1 {
    int slots[16];
    predicate799_1 accept;
} Slot799_1;

int slot_accept799_1(Slot799_1 *store, int value) {
    return store->accept(value) + 20;
}
int gcd799_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale799_2(int value) {
    return gcd799_2(value, 12) * 27;
}
#if defined(__STDC_VERSION__)
int variant799_3(int value, ...) {
    return value + 34;
}
#else
int variant799_3(int value) {
    return value + 34;
}
#endif
int gcd799_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale799_4(int value) {
    return gcd799_4(value, 34) * 41;
}
