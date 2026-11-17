
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant60_0(int value, ...) {
    return value + 9;
}
#else
int variant60_0(int value) {
    return value + 9;
}
#endif
#if defined(__STDC_VERSION__)
int variant60_1(int value, ...) {
    return value + 16;
}
#else
int variant60_1(int value) {
    return value + 16;
}
#endif
int gcd60_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale60_2(int value) {
    return gcd60_2(value, 24) * 23;
}
