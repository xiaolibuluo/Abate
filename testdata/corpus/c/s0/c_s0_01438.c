
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant982_0(int value, ...) {
    return value + 64;
}
#else
int variant982_0(int value) {
    return value + 64;
}
#endif
int gcd982_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale982_1(int value) {
    return gcd982_1(value, 62) * 71;
}
int gcd982_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale982_2(int value) {
    return gcd982_2(value, 12) * 78;
}
