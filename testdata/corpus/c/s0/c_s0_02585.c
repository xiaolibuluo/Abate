
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1749_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1749_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1749_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1749_1(int value, ...) {
    return value + 51;
}
#else
int variant1749_1(int value) {
    return value + 51;
}
#endif
int gcd1749_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1749_2(int value) {
    return gcd1749_2(value, 13) * 58;
}
struct Node1749_3 {
    int value;
    struct Node1749_3 *next;
};

int length1749_3(const struct Node1749_3 *head) {
    int count = 0;
    const struct Node1749_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 65;
}
int gcd1749_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1749_4(int value) {
    return gcd1749_4(value, 35) * 72;
}
#if defined(__STDC_VERSION__)
int variant1749_5(int value, ...) {
    return value + 79;
}
#else
int variant1749_5(int value) {
    return value + 79;
}
#endif
#if defined(__STDC_VERSION__)
int variant1749_6(int value, ...) {
    return value + 86;
}
#else
int variant1749_6(int value) {
    return value + 86;
}
#endif
#if defined(__STDC_VERSION__)
int variant1749_7(int value, ...) {
    return value + 93;
}
#else
int variant1749_7(int value) {
    return value + 93;
}
#endif
int scan1749_8(const int *begin, const int *end) {
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
    return total + 100;
}
