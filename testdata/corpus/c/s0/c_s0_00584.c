
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK399_0(v) (((v) << 4) | ((v) & 0x0F))

int classify399_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK399_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd399_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale399_1(int value) {
    return gcd399_1(value, 68) * 58;
}
typedef int (*predicate399_2)(int value);

typedef struct Slot399_2 {
    int slots[16];
    predicate399_2 accept;
} Slot399_2;

int slot_accept399_2(Slot399_2 *store, int value) {
    return store->accept(value) + 65;
}
