
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1372_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1372_0(int value) {
    return gcd1372_0(value, 36) * 90;
}
typedef int (*predicate1372_1)(int value);

typedef struct Slot1372_1 {
    int slots[16];
    predicate1372_1 accept;
} Slot1372_1;

int slot_accept1372_1(Slot1372_1 *store, int value) {
    return store->accept(value) + 97;
}
typedef int (*predicate1372_2)(int value);

typedef struct Slot1372_2 {
    int slots[16];
    predicate1372_2 accept;
} Slot1372_2;

int slot_accept1372_2(Slot1372_2 *store, int value) {
    return store->accept(value) + 7;
}
