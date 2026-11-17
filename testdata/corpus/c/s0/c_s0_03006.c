
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2035_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2035_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2035_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2035_1(const int *begin, const int *end) {
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
    return total + 83;
}
#define MASK2035_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2035_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2035_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2035_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2035_3(int value) {
    return gcd2035_3(value, 13) * 97;
}
int gcd2035_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2035_4(int value) {
    return gcd2035_4(value, 24) * 7;
}
#if defined(__STDC_VERSION__)
int variant2035_5(int value, ...) {
    return value + 14;
}
#else
int variant2035_5(int value) {
    return value + 14;
}
#endif
#if defined(__STDC_VERSION__)
int variant2035_6(int value, ...) {
    return value + 21;
}
#else
int variant2035_6(int value) {
    return value + 21;
}
#endif
struct Node2035_7 {
    int value;
    struct Node2035_7 *next;
};

int length2035_7(const struct Node2035_7 *head) {
    int count = 0;
    const struct Node2035_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
