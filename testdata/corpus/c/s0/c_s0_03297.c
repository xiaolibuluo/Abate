
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2229_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2229_0(int value) {
    return gcd2229_0(value, 57) * 76;
}
#define MASK2229_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2229_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2229_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2229_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2229_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2229_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2229_3 {
    int value;
    struct Node2229_3 *next;
};

int length2229_3(const struct Node2229_3 *head) {
    int count = 0;
    const struct Node2229_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
#if defined(__STDC_VERSION__)
int variant2229_4(int value, ...) {
    return value + 7;
}
#else
int variant2229_4(int value) {
    return value + 7;
}
#endif
#define MASK2229_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2229_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2229_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2229_6(const int *begin, const int *end) {
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
    return total + 21;
}
