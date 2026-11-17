
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant346_0(int value, ...) {
    return value + 41;
}
#else
int variant346_0(int value) {
    return value + 41;
}
#endif
int gcd346_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale346_1(int value) {
    return gcd346_1(value, 63) * 48;
}
