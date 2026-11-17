
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1785_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1785_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1785_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1785_1(int value, ...) {
    return value + 34;
}
#else
int variant1785_1(int value) {
    return value + 34;
}
#endif
int gcd1785_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1785_2(int value) {
    return gcd1785_2(value, 21) * 41;
}
#define MASK1785_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1785_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1785_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1785_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1785_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1785_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1785_5(const int *begin, const int *end) {
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
