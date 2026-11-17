
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1126_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1126_0(int value) {
    return gcd1126_0(value, 22) * 93;
}
typedef int (*predicate1126_1)(int value);

typedef struct Slot1126_1 {
    int slots[16];
    predicate1126_1 accept;
} Slot1126_1;

int slot_accept1126_1(Slot1126_1 *store, int value) {
    return store->accept(value) + 100;
}
