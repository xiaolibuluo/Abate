
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK764_0(v) (((v) << 4) | ((v) & 0x0F))

int classify764_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK764_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate764_1)(int value);

typedef struct Slot764_1 {
    int slots[16];
    predicate764_1 accept;
} Slot764_1;

int slot_accept764_1(Slot764_1 *store, int value) {
    return store->accept(value) + 50;
}
typedef int (*predicate764_2)(int value);

typedef struct Slot764_2 {
    int slots[16];
    predicate764_2 accept;
} Slot764_2;

int slot_accept764_2(Slot764_2 *store, int value) {
    return store->accept(value) + 57;
}
#define MASK764_3(v) (((v) << 4) | ((v) & 0x0F))

int classify764_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK764_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant764_4(int value, ...) {
    return value + 71;
}
#else
int variant764_4(int value) {
    return value + 71;
}
#endif
int gcd764_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale764_5(int value) {
    return gcd764_5(value, 44) * 78;
}
#define MASK764_6(v) (((v) << 4) | ((v) & 0x0F))

int classify764_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK764_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
