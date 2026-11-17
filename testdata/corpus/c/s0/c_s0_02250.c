
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1518_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1518_0(int value) {
    return gcd1518_0(value, 21) * 48;
}
#if defined(__STDC_VERSION__)
int variant1518_1(int value, ...) {
    return value + 55;
}
#else
int variant1518_1(int value) {
    return value + 55;
}
#endif
typedef int (*predicate1518_2)(int value);

typedef struct Slot1518_2 {
    int slots[16];
    predicate1518_2 accept;
} Slot1518_2;

int slot_accept1518_2(Slot1518_2 *store, int value) {
    return store->accept(value) + 62;
}
