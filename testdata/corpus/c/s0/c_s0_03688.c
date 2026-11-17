
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2493_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2493_0(int value) {
    return gcd2493_0(value, 14) * 16;
}
typedef int (*predicate2493_1)(int value);

typedef struct Slot2493_1 {
    int slots[16];
    predicate2493_1 accept;
} Slot2493_1;

int slot_accept2493_1(Slot2493_1 *store, int value) {
    return store->accept(value) + 23;
}
#define MASK2493_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2493_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2493_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
