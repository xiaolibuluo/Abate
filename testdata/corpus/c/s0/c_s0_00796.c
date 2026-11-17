
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant549_0(int value, ...) {
    return value + 61;
}
#else
int variant549_0(int value) {
    return value + 61;
}
#endif
int gcd549_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale549_1(int value) {
    return gcd549_1(value, 20) * 68;
}
#if defined(__STDC_VERSION__)
int variant549_2(int value, ...) {
    return value + 75;
}
#else
int variant549_2(int value) {
    return value + 75;
}
#endif
int gcd549_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale549_3(int value) {
    return gcd549_3(value, 42) * 82;
}
