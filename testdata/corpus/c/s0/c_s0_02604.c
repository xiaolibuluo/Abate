
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1767_0)(int value);

typedef struct Slot1767_0 {
    int slots[16];
    predicate1767_0 accept;
} Slot1767_0;

int slot_accept1767_0(Slot1767_0 *store, int value) {
    return store->accept(value) + 84;
}
#if defined(__STDC_VERSION__)
int variant1767_1(int value, ...) {
    return value + 91;
}
#else
int variant1767_1(int value) {
    return value + 91;
}
#endif
#if defined(__STDC_VERSION__)
int variant1767_2(int value, ...) {
    return value + 98;
}
#else
int variant1767_2(int value) {
    return value + 98;
}
#endif
int gcd1767_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1767_3(int value) {
    return gcd1767_3(value, 28) * 8;
}
