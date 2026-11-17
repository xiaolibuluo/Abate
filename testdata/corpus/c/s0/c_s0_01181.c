
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan809_0(const int *begin, const int *end) {
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
    return total + 46;
}
int gcd809_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale809_1(int value) {
    return gcd809_1(value, 10) * 53;
}
#if defined(__STDC_VERSION__)
int variant809_2(int value, ...) {
    return value + 60;
}
#else
int variant809_2(int value) {
    return value + 60;
}
#endif
int gcd809_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale809_3(int value) {
    return gcd809_3(value, 32) * 67;
}
#define MASK809_4(v) (((v) << 4) | ((v) & 0x0F))

int classify809_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK809_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan809_5(const int *begin, const int *end) {
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
    return total + 81;
}
