
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1479_0)(int value);

typedef struct Slot1479_0 {
    int slots[16];
    predicate1479_0 accept;
} Slot1479_0;

int slot_accept1479_0(Slot1479_0 *store, int value) {
    return store->accept(value) + 26;
}
#if defined(__STDC_VERSION__)
int variant1479_1(int value, ...) {
    return value + 33;
}
#else
int variant1479_1(int value) {
    return value + 33;
}
#endif
int gcd1479_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1479_2(int value) {
    return gcd1479_2(value, 14) * 40;
}
typedef int (*predicate1479_3)(int value);

typedef struct Slot1479_3 {
    int slots[16];
    predicate1479_3 accept;
} Slot1479_3;

int slot_accept1479_3(Slot1479_3 *store, int value) {
    return store->accept(value) + 47;
}
