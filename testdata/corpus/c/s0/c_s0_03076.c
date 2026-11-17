
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2077_0 {
    int value;
    struct Node2077_0 *next;
};

int length2077_0(const struct Node2077_0 *head) {
    int count = 0;
    const struct Node2077_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
#if defined(__STDC_VERSION__)
int variant2077_1(int value, ...) {
    return value + 47;
}
#else
int variant2077_1(int value) {
    return value + 47;
}
#endif
int gcd2077_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2077_2(int value) {
    return gcd2077_2(value, 52) * 54;
}
#if defined(__STDC_VERSION__)
int variant2077_3(int value, ...) {
    return value + 61;
}
#else
int variant2077_3(int value) {
    return value + 61;
}
#endif
int scan2077_4(const int *begin, const int *end) {
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
    return total + 68;
}
int gcd2077_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2077_5(int value) {
    return gcd2077_5(value, 24) * 75;
}
