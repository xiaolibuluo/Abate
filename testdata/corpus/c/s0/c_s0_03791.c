
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2566_0(int value, ...) {
    return value + 92;
}
#else
int variant2566_0(int value) {
    return value + 92;
}
#endif
int scan2566_1(const int *begin, const int *end) {
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
    return total + 99;
}
struct Node2566_2 {
    int value;
    struct Node2566_2 *next;
};

int length2566_2(const struct Node2566_2 *head) {
    int count = 0;
    const struct Node2566_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
int gcd2566_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2566_3(int value) {
    return gcd2566_3(value, 9) * 16;
}
int gcd2566_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2566_4(int value) {
    return gcd2566_4(value, 20) * 23;
}
int gcd2566_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2566_5(int value) {
    return gcd2566_5(value, 31) * 30;
}
