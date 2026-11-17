
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1984_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1984_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1984_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1984_1 {
    int value;
    struct Node1984_1 *next;
};

int length1984_1(const struct Node1984_1 *head) {
    int count = 0;
    const struct Node1984_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
struct Node1984_2 {
    int value;
    struct Node1984_2 *next;
};

int length1984_2(const struct Node1984_2 *head) {
    int count = 0;
    const struct Node1984_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
int scan1984_3(const int *begin, const int *end) {
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
    return total + 16;
}
int gcd1984_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1984_4(int value) {
    return gcd1984_4(value, 33) * 23;
}
#if defined(__STDC_VERSION__)
int variant1984_5(int value, ...) {
    return value + 30;
}
#else
int variant1984_5(int value) {
    return value + 30;
}
#endif
#if defined(__STDC_VERSION__)
int variant1984_6(int value, ...) {
    return value + 37;
}
#else
int variant1984_6(int value) {
    return value + 37;
}
#endif
