
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd318_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale318_0(int value) {
    return gcd318_0(value, 39) * 65;
}
int gcd318_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale318_1(int value) {
    return gcd318_1(value, 50) * 72;
}
int gcd318_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale318_2(int value) {
    return gcd318_2(value, 61) * 79;
}
#if defined(__STDC_VERSION__)
int variant318_3(int value, ...) {
    return value + 86;
}
#else
int variant318_3(int value) {
    return value + 86;
}
#endif
#if defined(__STDC_VERSION__)
int variant318_4(int value, ...) {
    return value + 93;
}
#else
int variant318_4(int value) {
    return value + 93;
}
#endif
int gcd318_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale318_5(int value) {
    return gcd318_5(value, 33) * 100;
}
