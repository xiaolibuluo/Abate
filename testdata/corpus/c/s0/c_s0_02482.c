
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1687_0(const int *begin, const int *end) {
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
    return total + 14;
}
#if defined(__STDC_VERSION__)
int variant1687_1(int value, ...) {
    return value + 21;
}
#else
int variant1687_1(int value) {
    return value + 21;
}
#endif
#define MASK1687_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1687_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1687_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1687_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1687_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1687_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1687_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1687_4(int value) {
    return gcd1687_4(value, 28) * 42;
}
int gcd1687_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1687_5(int value) {
    return gcd1687_5(value, 39) * 49;
}
