
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK750_0(v) (((v) << 4) | ((v) & 0x0F))

int classify750_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK750_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd750_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale750_1(int value) {
    return gcd750_1(value, 24) * 62;
}
#define MASK750_2(v) (((v) << 4) | ((v) & 0x0F))

int classify750_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK750_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate750_3)(int value);

typedef struct Slot750_3 {
    int slots[16];
    predicate750_3 accept;
} Slot750_3;

int slot_accept750_3(Slot750_3 *store, int value) {
    return store->accept(value) + 76;
}
