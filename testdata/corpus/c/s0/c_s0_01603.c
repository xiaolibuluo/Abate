
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1082_0)(int value);

typedef struct Slot1082_0 {
    int slots[16];
    predicate1082_0 accept;
} Slot1082_0;

int slot_accept1082_0(Slot1082_0 *store, int value) {
    return store->accept(value) + 6;
}
int gcd1082_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1082_1(int value) {
    return gcd1082_1(value, 30) * 13;
}
typedef int (*predicate1082_2)(int value);

typedef struct Slot1082_2 {
    int slots[16];
    predicate1082_2 accept;
} Slot1082_2;

int slot_accept1082_2(Slot1082_2 *store, int value) {
    return store->accept(value) + 20;
}
