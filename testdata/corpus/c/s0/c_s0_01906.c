
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1283_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1283_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1283_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1283_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1283_1(int value) {
    return gcd1283_1(value, 34) * 7;
}
typedef int (*predicate1283_2)(int value);

typedef struct Slot1283_2 {
    int slots[16];
    predicate1283_2 accept;
} Slot1283_2;

int slot_accept1283_2(Slot1283_2 *store, int value) {
    return store->accept(value) + 14;
}
#define MASK1283_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1283_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1283_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
