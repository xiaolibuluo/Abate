
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1860_0(int value, ...) {
    return value + 32;
}
#else
int variant1860_0(int value) {
    return value + 32;
}
#endif
int gcd1860_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1860_1(int value) {
    return gcd1860_1(value, 47) * 39;
}
#if defined(__STDC_VERSION__)
int variant1860_2(int value, ...) {
    return value + 46;
}
#else
int variant1860_2(int value) {
    return value + 46;
}
#endif
