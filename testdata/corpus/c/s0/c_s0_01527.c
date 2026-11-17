
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1035_0)(int value);

typedef struct Slot1035_0 {
    int slots[16];
    predicate1035_0 accept;
} Slot1035_0;

int slot_accept1035_0(Slot1035_0 *store, int value) {
    return store->accept(value) + 74;
}
int gcd1035_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1035_1(int value) {
    return gcd1035_1(value, 67) * 81;
}
#define MASK1035_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1035_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1035_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
