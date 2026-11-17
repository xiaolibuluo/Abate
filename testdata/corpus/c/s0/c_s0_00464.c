
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate316_0)(int value);

typedef struct Slot316_0 {
    int slots[16];
    predicate316_0 accept;
} Slot316_0;

int slot_accept316_0(Slot316_0 *store, int value) {
    return store->accept(value) + 39;
}
typedef int (*predicate316_1)(int value);

typedef struct Slot316_1 {
    int slots[16];
    predicate316_1 accept;
} Slot316_1;

int slot_accept316_1(Slot316_1 *store, int value) {
    return store->accept(value) + 46;
}
int gcd316_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale316_2(int value) {
    return gcd316_2(value, 47) * 53;
}
typedef int (*predicate316_3)(int value);

typedef struct Slot316_3 {
    int slots[16];
    predicate316_3 accept;
} Slot316_3;

int slot_accept316_3(Slot316_3 *store, int value) {
    return store->accept(value) + 60;
}
#define MASK316_4(v) (((v) << 4) | ((v) & 0x0F))

int classify316_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK316_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
