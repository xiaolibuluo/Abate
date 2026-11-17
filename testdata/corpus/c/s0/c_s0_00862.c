
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd593_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale593_0(int value) {
    return gcd593_0(value, 12) * 51;
}
#if defined(__STDC_VERSION__)
int variant593_1(int value, ...) {
    return value + 58;
}
#else
int variant593_1(int value) {
    return value + 58;
}
#endif
#if defined(__STDC_VERSION__)
int variant593_2(int value, ...) {
    return value + 65;
}
#else
int variant593_2(int value) {
    return value + 65;
}
#endif
