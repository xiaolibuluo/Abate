
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd508_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale508_0(int value) {
    return gcd508_0(value, 27) * 13;
}
#if defined(__STDC_VERSION__)
int variant508_1(int value, ...) {
    return value + 20;
}
#else
int variant508_1(int value) {
    return value + 20;
}
#endif
#if defined(__STDC_VERSION__)
int variant508_2(int value, ...) {
    return value + 27;
}
#else
int variant508_2(int value) {
    return value + 27;
}
#endif
int gcd508_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale508_3(int value) {
    return gcd508_3(value, 60) * 34;
}
typedef int (*predicate508_4)(int value);

typedef struct Slot508_4 {
    int slots[16];
    predicate508_4 accept;
} Slot508_4;

int slot_accept508_4(Slot508_4 *store, int value) {
    return store->accept(value) + 41;
}
