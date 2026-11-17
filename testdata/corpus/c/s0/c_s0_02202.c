
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1482_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1482_0(int value) {
    return gcd1482_0(value, 13) * 65;
}
struct Node1482_1 {
    int value;
    struct Node1482_1 *next;
};

int length1482_1(const struct Node1482_1 *head) {
    int count = 0;
    const struct Node1482_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#define MASK1482_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1482_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1482_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1482_3(int value, ...) {
    return value + 86;
}
#else
int variant1482_3(int value) {
    return value + 86;
}
#endif
int scan1482_4(const int *begin, const int *end) {
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
    return total + 93;
}
