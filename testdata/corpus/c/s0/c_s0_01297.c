
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd888_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale888_0(int value) {
    return gcd888_0(value, 64) * 6;
}
typedef int (*predicate888_1)(int value);

typedef struct Slot888_1 {
    int slots[16];
    predicate888_1 accept;
} Slot888_1;

int slot_accept888_1(Slot888_1 *store, int value) {
    return store->accept(value) + 13;
}
typedef int (*predicate888_2)(int value);

typedef struct Slot888_2 {
    int slots[16];
    predicate888_2 accept;
} Slot888_2;

int slot_accept888_2(Slot888_2 *store, int value) {
    return store->accept(value) + 20;
}
