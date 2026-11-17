
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1565_0(int value, ...) {
    return value + 77;
}
#else
int variant1565_0(int value) {
    return value + 77;
}
#endif
int gcd1565_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1565_1(int value) {
    return gcd1565_1(value, 56) * 84;
}
typedef int (*predicate1565_2)(int value);

typedef struct Slot1565_2 {
    int slots[16];
    predicate1565_2 accept;
} Slot1565_2;

int slot_accept1565_2(Slot1565_2 *store, int value) {
    return store->accept(value) + 91;
}
