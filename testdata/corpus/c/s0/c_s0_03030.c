
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2049_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2049_0(int value) {
    return gcd2049_0(value, 17) * 64;
}
#if defined(__STDC_VERSION__)
int variant2049_1(int value, ...) {
    return value + 71;
}
#else
int variant2049_1(int value) {
    return value + 71;
}
#endif
