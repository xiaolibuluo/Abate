
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1244_0(int value, ...) {
    return value + 75;
}
#else
int variant1244_0(int value) {
    return value + 75;
}
#endif
int scan1244_1(const int *begin, const int *end) {
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
    return total + 82;
}
int gcd1244_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1244_2(int value) {
    return gcd1244_2(value, 16) * 89;
}
struct Node1244_3 {
    int value;
    struct Node1244_3 *next;
};

int length1244_3(const struct Node1244_3 *head) {
    int count = 0;
    const struct Node1244_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
#define MASK1244_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1244_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1244_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
