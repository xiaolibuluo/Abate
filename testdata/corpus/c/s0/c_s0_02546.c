
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1730_0)(int value);

typedef struct Slot1730_0 {
    int slots[16];
    predicate1730_0 accept;
} Slot1730_0;

int slot_accept1730_0(Slot1730_0 *store, int value) {
    return store->accept(value) + 88;
}
int gcd1730_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1730_1(int value) {
    return gcd1730_1(value, 52) * 95;
}
int gcd1730_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1730_2(int value) {
    return gcd1730_2(value, 63) * 5;
}
#define MASK1730_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1730_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1730_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1730_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1730_4(int value) {
    return gcd1730_4(value, 24) * 19;
}
