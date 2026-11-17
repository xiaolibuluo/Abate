
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2406_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2406_0(int value) {
    return gcd2406_0(value, 15) * 49;
}
#define MASK2406_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2406_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2406_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2406_2 {
    int value;
    struct Node2406_2 *next;
};

int length2406_2(const struct Node2406_2 *head) {
    int count = 0;
    const struct Node2406_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 63;
}
#if defined(__STDC_VERSION__)
int variant2406_3(int value, ...) {
    return value + 70;
}
#else
int variant2406_3(int value) {
    return value + 70;
}
#endif
int scan2406_4(const int *begin, const int *end) {
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
    return total + 77;
}
