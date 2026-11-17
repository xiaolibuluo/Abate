
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1815_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1815_0(int value) {
    return gcd1815_0(value, 26) * 29;
}
typedef int (*predicate1815_1)(int value);

typedef struct Slot1815_1 {
    int slots[16];
    predicate1815_1 accept;
} Slot1815_1;

int slot_accept1815_1(Slot1815_1 *store, int value) {
    return store->accept(value) + 36;
}
typedef int (*predicate1815_2)(int value);

typedef struct Slot1815_2 {
    int slots[16];
    predicate1815_2 accept;
} Slot1815_2;

int slot_accept1815_2(Slot1815_2 *store, int value) {
    return store->accept(value) + 43;
}
