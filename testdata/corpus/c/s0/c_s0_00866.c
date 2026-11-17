
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd596_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale596_0(int value) {
    return gcd596_0(value, 33) * 90;
}
#if defined(__STDC_VERSION__)
int variant596_1(int value, ...) {
    return value + 97;
}
#else
int variant596_1(int value) {
    return value + 97;
}
#endif
int gcd596_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale596_2(int value) {
    return gcd596_2(value, 55) * 7;
}
