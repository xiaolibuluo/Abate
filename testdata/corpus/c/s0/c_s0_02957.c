
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2000_0)(int value);

typedef struct Slot2000_0 {
    int slots[16];
    predicate2000_0 accept;
} Slot2000_0;

int slot_accept2000_0(Slot2000_0 *store, int value) {
    return store->accept(value) + 9;
}
int gcd2000_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2000_1(int value) {
    return gcd2000_1(value, 51) * 16;
}
