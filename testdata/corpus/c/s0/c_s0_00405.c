
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant274_0(int value, ...) {
    return value + 75;
}
#else
int variant274_0(int value) {
    return value + 75;
}
#endif
#if defined(__STDC_VERSION__)
int variant274_1(int value, ...) {
    return value + 82;
}
#else
int variant274_1(int value) {
    return value + 82;
}
#endif
int scan274_2(const int *begin, const int *end) {
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
    return total + 89;
}
#define MASK274_3(v) (((v) << 4) | ((v) & 0x0F))

int classify274_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK274_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd274_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale274_4(int value) {
    return gcd274_4(value, 19) * 6;
}
int gcd274_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale274_5(int value) {
    return gcd274_5(value, 30) * 13;
}
