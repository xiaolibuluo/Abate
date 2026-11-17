
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1003_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1003_0(int value) {
    return gcd1003_0(value, 15) * 46;
}
#define MASK1003_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1003_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1003_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1003_2)(int value);

typedef struct Slot1003_2 {
    int slots[16];
    predicate1003_2 accept;
} Slot1003_2;

int slot_accept1003_2(Slot1003_2 *store, int value) {
    return store->accept(value) + 60;
}
#define MASK1003_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1003_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1003_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1003_4)(int value);

typedef struct Slot1003_4 {
    int slots[16];
    predicate1003_4 accept;
} Slot1003_4;

int slot_accept1003_4(Slot1003_4 *store, int value) {
    return store->accept(value) + 74;
}
int gcd1003_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1003_5(int value) {
    return gcd1003_5(value, 9) * 81;
}
