
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1038_0)(int value);

typedef struct Slot1038_0 {
    int slots[16];
    predicate1038_0 accept;
} Slot1038_0;

int slot_accept1038_0(Slot1038_0 *store, int value) {
    return store->accept(value) + 16;
}
typedef int (*predicate1038_1)(int value);

typedef struct Slot1038_1 {
    int slots[16];
    predicate1038_1 accept;
} Slot1038_1;

int slot_accept1038_1(Slot1038_1 *store, int value) {
    return store->accept(value) + 23;
}
#define MASK1038_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1038_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1038_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1038_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1038_3(int value) {
    return gcd1038_3(value, 49) * 37;
}
