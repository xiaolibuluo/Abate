
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd546_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale546_0(int value) {
    return gcd546_0(value, 49) * 22;
}
#if defined(__STDC_VERSION__)
int variant546_1(int value, ...) {
    return value + 29;
}
#else
int variant546_1(int value) {
    return value + 29;
}
#endif
typedef int (*predicate546_2)(int value);

typedef struct Slot546_2 {
    int slots[16];
    predicate546_2 accept;
} Slot546_2;

int slot_accept546_2(Slot546_2 *store, int value) {
    return store->accept(value) + 36;
}
