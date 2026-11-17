
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan241_0(const int *begin, const int *end) {
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
    return total + 34;
}
int scan241_1(const int *begin, const int *end) {
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
    return total + 41;
}
#if defined(__STDC_VERSION__)
int variant241_2(int value, ...) {
    return value + 48;
}
#else
int variant241_2(int value) {
    return value + 48;
}
#endif
int gcd241_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale241_3(int value) {
    return gcd241_3(value, 21) * 55;
}
int scan241_4(const int *begin, const int *end) {
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
    return total + 62;
}
#define MASK241_5(v) (((v) << 4) | ((v) & 0x0F))

int classify241_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK241_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant241_6(int value, ...) {
    return value + 76;
}
#else
int variant241_6(int value) {
    return value + 76;
}
#endif
#if defined(__STDC_VERSION__)
int variant241_7(int value, ...) {
    return value + 83;
}
#else
int variant241_7(int value) {
    return value + 83;
}
#endif
