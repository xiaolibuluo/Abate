
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1174_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1174_0(int value) {
    return gcd1174_0(value, 53) * 38;
}
typedef int (*predicate1174_1)(int value);

typedef struct Slot1174_1 {
    int slots[16];
    predicate1174_1 accept;
} Slot1174_1;

int slot_accept1174_1(Slot1174_1 *store, int value) {
    return store->accept(value) + 45;
}
