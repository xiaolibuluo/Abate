
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant119_0(int value, ...) {
    return value + 97;
}
#else
int variant119_0(int value) {
    return value + 97;
}
#endif
#define MASK119_1(v) (((v) << 4) | ((v) & 0x0F))

int classify119_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK119_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd119_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale119_2(int value) {
    return gcd119_2(value, 10) * 14;
}
typedef int (*predicate119_3)(int value);

typedef struct Slot119_3 {
    int slots[16];
    predicate119_3 accept;
} Slot119_3;

int slot_accept119_3(Slot119_3 *store, int value) {
    return store->accept(value) + 21;
}
#define MASK119_4(v) (((v) << 4) | ((v) & 0x0F))

int classify119_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK119_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
