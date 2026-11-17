
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2588_0)(int value);

typedef struct Slot2588_0 {
    int slots[16];
    predicate2588_0 accept;
} Slot2588_0;

int slot_accept2588_0(Slot2588_0 *store, int value) {
    return store->accept(value) + 87;
}
int gcd2588_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2588_1(int value) {
    return gcd2588_1(value, 19) * 94;
}
#if defined(__STDC_VERSION__)
int variant2588_2(int value, ...) {
    return value + 101;
}
#else
int variant2588_2(int value) {
    return value + 101;
}
#endif
