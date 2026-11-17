
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2099_0(int value, ...) {
    return value + 35;
}
#else
int variant2099_0(int value) {
    return value + 35;
}
#endif
int gcd2099_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2099_1(int value) {
    return gcd2099_1(value, 12) * 42;
}
