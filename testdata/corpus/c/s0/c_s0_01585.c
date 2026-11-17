
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1070_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1070_0(int value) {
    return gcd1070_0(value, 57) * 44;
}
int gcd1070_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1070_1(int value) {
    return gcd1070_1(value, 68) * 51;
}
#if defined(__STDC_VERSION__)
int variant1070_2(int value, ...) {
    return value + 58;
}
#else
int variant1070_2(int value) {
    return value + 58;
}
#endif
int gcd1070_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1070_3(int value) {
    return gcd1070_3(value, 29) * 65;
}
int scan1070_4(const int *begin, const int *end) {
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
    return total + 72;
}
int scan1070_5(const int *begin, const int *end) {
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
struct Node1070_6 {
    int value;
    struct Node1070_6 *next;
};

int length1070_6(const struct Node1070_6 *head) {
    int count = 0;
    const struct Node1070_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
#define MASK1070_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1070_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1070_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
