
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2342_0(int value, ...) {
    return value + 90;
}
#else
int variant2342_0(int value) {
    return value + 90;
}
#endif
int gcd2342_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2342_1(int value) {
    return gcd2342_1(value, 66) * 97;
}
