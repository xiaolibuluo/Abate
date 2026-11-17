
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1385_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1385_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1385_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1385_1)(int value);

typedef struct Slot1385_1 {
    int slots[16];
    predicate1385_1 accept;
} Slot1385_1;

int slot_accept1385_1(Slot1385_1 *store, int value) {
    return store->accept(value) + 72;
}
typedef int (*predicate1385_2)(int value);

typedef struct Slot1385_2 {
    int slots[16];
    predicate1385_2 accept;
} Slot1385_2;

int slot_accept1385_2(Slot1385_2 *store, int value) {
    return store->accept(value) + 79;
}
int gcd1385_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1385_3(int value) {
    return gcd1385_3(value, 38) * 86;
}
typedef int (*predicate1385_4)(int value);

typedef struct Slot1385_4 {
    int slots[16];
    predicate1385_4 accept;
} Slot1385_4;

int slot_accept1385_4(Slot1385_4 *store, int value) {
    return store->accept(value) + 93;
}
