
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2627_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2627_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2627_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2627_1(const int *begin, const int *end) {
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
    return total + 19;
}
int scan2627_2(const int *begin, const int *end) {
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
    return total + 26;
}
#if defined(__STDC_VERSION__)
int variant2627_3(int value, ...) {
    return value + 33;
}
#else
int variant2627_3(int value) {
    return value + 33;
}
#endif
struct Node2627_4 {
    int value;
    struct Node2627_4 *next;
};

int length2627_4(const struct Node2627_4 *head) {
    int count = 0;
    const struct Node2627_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
#if defined(__STDC_VERSION__)
int variant2627_5(int value, ...) {
    return value + 47;
}
#else
int variant2627_5(int value) {
    return value + 47;
}
#endif
int gcd2627_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2627_6(int value) {
    return gcd2627_6(value, 42) * 54;
}
int scan2627_7(const int *begin, const int *end) {
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
    return total + 61;
}
