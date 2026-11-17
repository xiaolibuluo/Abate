
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd254_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale254_0(int value) {
    return gcd254_0(value, 18) * 9;
}
typedef int (*predicate254_1)(int value);

typedef struct Slot254_1 {
    int slots[16];
    predicate254_1 accept;
} Slot254_1;

int slot_accept254_1(Slot254_1 *store, int value) {
    return store->accept(value) + 16;
}
typedef int (*predicate254_2)(int value);

typedef struct Slot254_2 {
    int slots[16];
    predicate254_2 accept;
} Slot254_2;

int slot_accept254_2(Slot254_2 *store, int value) {
    return store->accept(value) + 23;
}
