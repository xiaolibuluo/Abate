
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1600_0(const int *begin, const int *end) {
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
    return total + 47;
}
int scan1600_1(const int *begin, const int *end) {
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
    return total + 54;
}
struct Node1600_2 {
    int value;
    struct Node1600_2 *next;
};

int length1600_2(const struct Node1600_2 *head) {
    int count = 0;
    const struct Node1600_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 61;
}
int gcd1600_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1600_3(int value) {
    return gcd1600_3(value, 18) * 68;
}
int gcd1600_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1600_4(int value) {
    return gcd1600_4(value, 29) * 75;
}
#define MASK1600_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1600_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1600_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1600_6(int value, ...) {
    return value + 89;
}
#else
int variant1600_6(int value) {
    return value + 89;
}
#endif
