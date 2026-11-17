
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2557_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2557_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2557_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2557_1)(int value);

typedef struct Slot2557_1 {
    int slots[16];
    predicate2557_1 accept;
} Slot2557_1;

int slot_accept2557_1(Slot2557_1 *store, int value) {
    return store->accept(value) + 79;
}
int gcd2557_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2557_2(int value) {
    return gcd2557_2(value, 57) * 86;
}
typedef int (*predicate2557_3)(int value);

typedef struct Slot2557_3 {
    int slots[16];
    predicate2557_3 accept;
} Slot2557_3;

int slot_accept2557_3(Slot2557_3 *store, int value) {
    return store->accept(value) + 93;
}
typedef int (*predicate2557_4)(int value);

typedef struct Slot2557_4 {
    int slots[16];
    predicate2557_4 accept;
} Slot2557_4;

int slot_accept2557_4(Slot2557_4 *store, int value) {
    return store->accept(value) + 100;
}
