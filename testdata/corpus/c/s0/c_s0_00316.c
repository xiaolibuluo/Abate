
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd218_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale218_0(int value) {
    return gcd218_0(value, 10) * 26;
}
typedef int (*predicate218_1)(int value);

typedef struct Slot218_1 {
    int slots[16];
    predicate218_1 accept;
} Slot218_1;

int slot_accept218_1(Slot218_1 *store, int value) {
    return store->accept(value) + 33;
}
int gcd218_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale218_2(int value) {
    return gcd218_2(value, 32) * 40;
}
typedef int (*predicate218_3)(int value);

typedef struct Slot218_3 {
    int slots[16];
    predicate218_3 accept;
} Slot218_3;

int slot_accept218_3(Slot218_3 *store, int value) {
    return store->accept(value) + 47;
}
#define MASK218_4(v) (((v) << 4) | ((v) & 0x0F))

int classify218_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK218_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
