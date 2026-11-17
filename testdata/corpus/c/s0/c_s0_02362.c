
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1593_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1593_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1593_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1593_1(int value, ...) {
    return value + 60;
}
#else
int variant1593_1(int value) {
    return value + 60;
}
#endif
int gcd1593_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1593_2(int value) {
    return gcd1593_2(value, 19) * 67;
}
int gcd1593_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1593_3(int value) {
    return gcd1593_3(value, 30) * 74;
}
int gcd1593_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1593_4(int value) {
    return gcd1593_4(value, 41) * 81;
}
#if defined(__STDC_VERSION__)
int variant1593_5(int value, ...) {
    return value + 88;
}
#else
int variant1593_5(int value) {
    return value + 88;
}
#endif
int scan1593_6(const int *begin, const int *end) {
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
