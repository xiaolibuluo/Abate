
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd899_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale899_0(int value) {
    return gcd899_0(value, 19) * 52;
}
#if defined(__STDC_VERSION__)
int variant899_1(int value, ...) {
    return value + 59;
}
#else
int variant899_1(int value) {
    return value + 59;
}
#endif
typedef int (*predicate899_2)(int value);

typedef struct Slot899_2 {
    int slots[16];
    predicate899_2 accept;
} Slot899_2;

int slot_accept899_2(Slot899_2 *store, int value) {
    return store->accept(value) + 66;
}
#if defined(__STDC_VERSION__)
int variant899_3(int value, ...) {
    return value + 73;
}
#else
int variant899_3(int value) {
    return value + 73;
}
#endif
#if defined(__STDC_VERSION__)
int variant899_4(int value, ...) {
    return value + 80;
}
#else
int variant899_4(int value) {
    return value + 80;
}
#endif
int gcd899_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale899_5(int value) {
    return gcd899_5(value, 13) * 87;
}
