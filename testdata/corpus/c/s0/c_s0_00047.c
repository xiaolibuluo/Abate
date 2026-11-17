
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant29_0(int value, ...) {
    return value + 91;
}
#else
int variant29_0(int value) {
    return value + 91;
}
#endif
int gcd29_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale29_1(int value) {
    return gcd29_1(value, 40) * 98;
}
int gcd29_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale29_2(int value) {
    return gcd29_2(value, 51) * 8;
}
int gcd29_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale29_3(int value) {
    return gcd29_3(value, 62) * 15;
}
int gcd29_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale29_4(int value) {
    return gcd29_4(value, 12) * 22;
}
int gcd29_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale29_5(int value) {
    return gcd29_5(value, 23) * 29;
}
