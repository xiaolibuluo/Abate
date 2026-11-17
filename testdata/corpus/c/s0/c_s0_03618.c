
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2444_0)(int value);

typedef struct Slot2444_0 {
    int slots[16];
    predicate2444_0 accept;
} Slot2444_0;

int slot_accept2444_0(Slot2444_0 *store, int value) {
    return store->accept(value) + 58;
}
int gcd2444_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2444_1(int value) {
    return gcd2444_1(value, 48) * 65;
}
#define MASK2444_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2444_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2444_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
