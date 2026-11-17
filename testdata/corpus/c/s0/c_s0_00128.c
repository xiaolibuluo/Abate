
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan87_0(const int *begin, const int *end) {
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
    return total + 69;
}
int scan87_1(const int *begin, const int *end) {
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
    return total + 76;
}
#define MASK87_2(v) (((v) << 4) | ((v) & 0x0F))

int classify87_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK87_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant87_3(int value, ...) {
    return value + 90;
}
#else
int variant87_3(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant87_4(int value, ...) {
    return value + 97;
}
#else
int variant87_4(int value) {
    return value + 97;
}
#endif
int gcd87_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale87_5(int value) {
    return gcd87_5(value, 63) * 7;
}
struct Node87_6 {
    int value;
    struct Node87_6 *next;
};

int length87_6(const struct Node87_6 *head) {
    int count = 0;
    const struct Node87_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
