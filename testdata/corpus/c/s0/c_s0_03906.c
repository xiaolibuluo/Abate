
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2645_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2645_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2645_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2645_1)(int value);

typedef struct Slot2645_1 {
    int slots[16];
    predicate2645_1 accept;
} Slot2645_1;

int slot_accept2645_1(Slot2645_1 *store, int value) {
    return store->accept(value) + 59;
}
int gcd2645_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2645_2(int value) {
    return gcd2645_2(value, 63) * 66;
}
int gcd2645_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2645_3(int value) {
    return gcd2645_3(value, 13) * 73;
}
