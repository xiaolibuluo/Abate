
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2585_0(int value, ...) {
    return value + 48;
}
#else
int variant2585_0(int value) {
    return value + 48;
}
#endif
#if defined(__STDC_VERSION__)
int variant2585_1(int value, ...) {
    return value + 55;
}
#else
int variant2585_1(int value) {
    return value + 55;
}
#endif
int gcd2585_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2585_2(int value) {
    return gcd2585_2(value, 9) * 62;
}
