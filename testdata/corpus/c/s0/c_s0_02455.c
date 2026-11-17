
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1667_0)(int value);

typedef struct Slot1667_0 {
    int slots[16];
    predicate1667_0 accept;
} Slot1667_0;

int slot_accept1667_0(Slot1667_0 *store, int value) {
    return store->accept(value) + 45;
}
#if defined(__STDC_VERSION__)
int variant1667_1(int value, ...) {
    return value + 52;
}
#else
int variant1667_1(int value) {
    return value + 52;
}
#endif
int gcd1667_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1667_2(int value) {
    return gcd1667_2(value, 49) * 59;
}
