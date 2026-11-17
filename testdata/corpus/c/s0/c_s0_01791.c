
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1201_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1201_0(int value) {
    return gcd1201_0(value, 59) * 98;
}
int gcd1201_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1201_1(int value) {
    return gcd1201_1(value, 9) * 8;
}
int gcd1201_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1201_2(int value) {
    return gcd1201_2(value, 20) * 15;
}
#if defined(__STDC_VERSION__)
int variant1201_3(int value, ...) {
    return value + 22;
}
#else
int variant1201_3(int value) {
    return value + 22;
}
#endif
#if defined(__STDC_VERSION__)
int variant1201_4(int value, ...) {
    return value + 29;
}
#else
int variant1201_4(int value) {
    return value + 29;
}
#endif
int scan1201_5(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 36;
}
#if defined(__STDC_VERSION__)
int variant1201_6(int value, ...) {
    return value + 43;
}
#else
int variant1201_6(int value) {
    return value + 43;
}
#endif
