
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd296_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale296_0(int value) {
    return gcd296_0(value, 68) * 70;
}
typedef int (*predicate296_1)(int value);

typedef struct Slot296_1 {
    int slots[16];
    predicate296_1 accept;
} Slot296_1;

int slot_accept296_1(Slot296_1 *store, int value) {
    return store->accept(value) + 77;
}
