
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1128_0(int value, ...) {
    return value + 22;
}
#else
int variant1128_0(int value) {
    return value + 22;
}
#endif
#if defined(__STDC_VERSION__)
int variant1128_1(int value, ...) {
    return value + 29;
}
#else
int variant1128_1(int value) {
    return value + 29;
}
#endif
int gcd1128_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1128_2(int value) {
    return gcd1128_2(value, 58) * 36;
}
