
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2481_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2481_0(int value) {
    return gcd2481_0(value, 52) * 54;
}
struct Node2481_1 {
    int value;
    struct Node2481_1 *next;
};

int length2481_1(const struct Node2481_1 *head) {
    int count = 0;
    const struct Node2481_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
int gcd2481_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2481_2(int value) {
    return gcd2481_2(value, 13) * 68;
}
int gcd2481_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2481_3(int value) {
    return gcd2481_3(value, 24) * 75;
}
#if defined(__STDC_VERSION__)
int variant2481_4(int value, ...) {
    return value + 82;
}
#else
int variant2481_4(int value) {
    return value + 82;
}
#endif
int gcd2481_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2481_5(int value) {
    return gcd2481_5(value, 46) * 89;
}
#define MASK2481_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2481_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2481_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2481_7(const int *begin, const int *end) {
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
    return total + 6;
}
