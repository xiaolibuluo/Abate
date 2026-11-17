
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd349_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale349_0(int value) {
    return gcd349_0(value, 12) * 80;
}
#if defined(__STDC_VERSION__)
int variant349_1(int value, ...) {
    return value + 87;
}
#else
int variant349_1(int value) {
    return value + 87;
}
#endif
typedef int (*predicate349_2)(int value);

typedef struct Slot349_2 {
    int slots[16];
    predicate349_2 accept;
} Slot349_2;

int slot_accept349_2(Slot349_2 *store, int value) {
    return store->accept(value) + 94;
}
