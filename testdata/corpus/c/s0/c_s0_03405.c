
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2297_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2297_0(int value) {
    return gcd2297_0(value, 45) * 87;
}
int gcd2297_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2297_1(int value) {
    return gcd2297_1(value, 56) * 94;
}
typedef int (*predicate2297_2)(int value);

typedef struct Slot2297_2 {
    int slots[16];
    predicate2297_2 accept;
} Slot2297_2;

int slot_accept2297_2(Slot2297_2 *store, int value) {
    return store->accept(value) + 101;
}
