
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd416_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale416_0(int value) {
    return gcd416_0(value, 54) * 78;
}
int scan416_1(const int *begin, const int *end) {
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
    return total + 85;
}
#if defined(__STDC_VERSION__)
int variant416_2(int value, ...) {
    return value + 92;
}
#else
int variant416_2(int value) {
    return value + 92;
}
#endif
#define MASK416_3(v) (((v) << 4) | ((v) & 0x0F))

int classify416_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK416_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant416_4(int value, ...) {
    return value + 9;
}
#else
int variant416_4(int value) {
    return value + 9;
}
#endif
#define MASK416_5(v) (((v) << 4) | ((v) & 0x0F))

int classify416_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK416_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
