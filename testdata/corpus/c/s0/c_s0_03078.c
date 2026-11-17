
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2079_0(const int *begin, const int *end) {
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
    return total + 66;
}
int scan2079_1(const int *begin, const int *end) {
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
    return total + 73;
}
struct Node2079_2 {
    int value;
    struct Node2079_2 *next;
};

int length2079_2(const struct Node2079_2 *head) {
    int count = 0;
    const struct Node2079_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
int gcd2079_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2079_3(int value) {
    return gcd2079_3(value, 16) * 87;
}
int gcd2079_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2079_4(int value) {
    return gcd2079_4(value, 27) * 94;
}
#if defined(__STDC_VERSION__)
int variant2079_5(int value, ...) {
    return value + 101;
}
#else
int variant2079_5(int value) {
    return value + 101;
}
#endif
