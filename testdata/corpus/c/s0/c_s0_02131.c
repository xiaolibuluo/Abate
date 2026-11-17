
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1436_0)(int value);

typedef struct Slot1436_0 {
    int slots[16];
    predicate1436_0 accept;
} Slot1436_0;

int slot_accept1436_0(Slot1436_0 *store, int value) {
    return store->accept(value) + 49;
}
int gcd1436_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1436_1(int value) {
    return gcd1436_1(value, 68) * 56;
}
typedef int (*predicate1436_2)(int value);

typedef struct Slot1436_2 {
    int slots[16];
    predicate1436_2 accept;
} Slot1436_2;

int slot_accept1436_2(Slot1436_2 *store, int value) {
    return store->accept(value) + 63;
}
#if defined(__STDC_VERSION__)
int variant1436_3(int value, ...) {
    return value + 70;
}
#else
int variant1436_3(int value) {
    return value + 70;
}
#endif
int gcd1436_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1436_4(int value) {
    return gcd1436_4(value, 40) * 77;
}
