
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd868_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale868_0(int value) {
    return gcd868_0(value, 46) * 37;
}
#if defined(__STDC_VERSION__)
int variant868_1(int value, ...) {
    return value + 44;
}
#else
int variant868_1(int value) {
    return value + 44;
}
#endif
struct Node868_2 {
    int value;
    struct Node868_2 *next;
};

int length868_2(const struct Node868_2 *head) {
    int count = 0;
    const struct Node868_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 51;
}
int scan868_3(const int *begin, const int *end) {
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
    return total + 58;
}
#define MASK868_4(v) (((v) << 4) | ((v) & 0x0F))

int classify868_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK868_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node868_5 {
    int value;
    struct Node868_5 *next;
};

int length868_5(const struct Node868_5 *head) {
    int count = 0;
    const struct Node868_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
int scan868_6(const int *begin, const int *end) {
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
    return total + 79;
}
