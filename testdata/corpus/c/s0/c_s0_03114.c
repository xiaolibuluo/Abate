
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2108_0(int value, ...) {
    return value + 55;
}
#else
int variant2108_0(int value) {
    return value + 55;
}
#endif
int gcd2108_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2108_1(int value) {
    return gcd2108_1(value, 14) * 62;
}
int gcd2108_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2108_2(int value) {
    return gcd2108_2(value, 25) * 69;
}
typedef int (*predicate2108_3)(int value);

typedef struct Slot2108_3 {
    int slots[16];
    predicate2108_3 accept;
} Slot2108_3;

int slot_accept2108_3(Slot2108_3 *store, int value) {
    return store->accept(value) + 76;
}
