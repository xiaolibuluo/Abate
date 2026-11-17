
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1427_0)(int value);

typedef struct Slot1427_0 {
    int slots[16];
    predicate1427_0 accept;
} Slot1427_0;

int slot_accept1427_0(Slot1427_0 *store, int value) {
    return store->accept(value) + 29;
}
#if defined(__STDC_VERSION__)
int variant1427_1(int value, ...) {
    return value + 36;
}
#else
int variant1427_1(int value) {
    return value + 36;
}
#endif
int gcd1427_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1427_2(int value) {
    return gcd1427_2(value, 16) * 43;
}
#if defined(__STDC_VERSION__)
int variant1427_3(int value, ...) {
    return value + 50;
}
#else
int variant1427_3(int value) {
    return value + 50;
}
#endif
