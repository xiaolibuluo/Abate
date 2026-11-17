
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2147_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2147_0(int value) {
    return gcd2147_0(value, 32) * 77;
}
typedef int (*predicate2147_1)(int value);

typedef struct Slot2147_1 {
    int slots[16];
    predicate2147_1 accept;
} Slot2147_1;

int slot_accept2147_1(Slot2147_1 *store, int value) {
    return store->accept(value) + 84;
}
