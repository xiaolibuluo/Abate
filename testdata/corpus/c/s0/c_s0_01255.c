
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant857_0(int value, ...) {
    return value + 88;
}
#else
int variant857_0(int value) {
    return value + 88;
}
#endif
int scan857_1(const int *begin, const int *end) {
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
    return total + 95;
}
int gcd857_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale857_2(int value) {
    return gcd857_2(value, 52) * 5;
}
int gcd857_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale857_3(int value) {
    return gcd857_3(value, 63) * 12;
}
#define MASK857_4(v) (((v) << 4) | ((v) & 0x0F))

int classify857_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK857_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK857_5(v) (((v) << 4) | ((v) & 0x0F))

int classify857_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK857_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
