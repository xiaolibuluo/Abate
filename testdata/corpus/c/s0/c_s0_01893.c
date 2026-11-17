
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1273_0)(int value);

typedef struct Slot1273_0 {
    int slots[16];
    predicate1273_0 accept;
} Slot1273_0;

int slot_accept1273_0(Slot1273_0 *store, int value) {
    return store->accept(value) + 64;
}
int gcd1273_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1273_1(int value) {
    return gcd1273_1(value, 25) * 71;
}
