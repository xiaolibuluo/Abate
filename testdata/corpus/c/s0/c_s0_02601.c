
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1764_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1764_0(int value) {
    return gcd1764_0(value, 35) * 45;
}
#define MASK1764_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1764_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1764_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1764_2)(int value);

typedef struct Slot1764_2 {
    int slots[16];
    predicate1764_2 accept;
} Slot1764_2;

int slot_accept1764_2(Slot1764_2 *store, int value) {
    return store->accept(value) + 59;
}
