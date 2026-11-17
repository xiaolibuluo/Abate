
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1806_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1806_0(int value) {
    return gcd1806_0(value, 24) * 9;
}
#if defined(__STDC_VERSION__)
int variant1806_1(int value, ...) {
    return value + 16;
}
#else
int variant1806_1(int value) {
    return value + 16;
}
#endif
