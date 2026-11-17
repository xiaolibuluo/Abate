
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1391_0 {
    int value;
    struct Node1391_0 *next;
};

int length1391_0(const struct Node1391_0 *head) {
    int count = 0;
    const struct Node1391_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
int gcd1391_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1391_1(int value) {
    return gcd1391_1(value, 58) * 53;
}
#if defined(__STDC_VERSION__)
int variant1391_2(int value, ...) {
    return value + 60;
}
#else
int variant1391_2(int value) {
    return value + 60;
}
#endif
#if defined(__STDC_VERSION__)
int variant1391_3(int value, ...) {
    return value + 67;
}
#else
int variant1391_3(int value) {
    return value + 67;
}
#endif
#define MASK1391_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1391_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1391_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1391_5(const int *begin, const int *end) {
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
    return total + 81;
}
