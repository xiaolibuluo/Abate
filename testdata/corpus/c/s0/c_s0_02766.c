
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1884_0(int value, ...) {
    return value + 53;
}
#else
int variant1884_0(int value) {
    return value + 53;
}
#endif
struct Node1884_1 {
    int value;
    struct Node1884_1 *next;
};

int length1884_1(const struct Node1884_1 *head) {
    int count = 0;
    const struct Node1884_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
int gcd1884_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1884_2(int value) {
    return gcd1884_2(value, 43) * 67;
}
int scan1884_3(const int *begin, const int *end) {
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
    return total + 74;
}
#if defined(__STDC_VERSION__)
int variant1884_4(int value, ...) {
    return value + 81;
}
#else
int variant1884_4(int value) {
    return value + 81;
}
#endif
int scan1884_5(const int *begin, const int *end) {
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
    return total + 88;
}
