
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1223_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1223_0(int value) {
    return gcd1223_0(value, 30) * 93;
}
#if defined(__STDC_VERSION__)
int variant1223_1(int value, ...) {
    return value + 100;
}
#else
int variant1223_1(int value) {
    return value + 100;
}
#endif
