
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node994_0 {
    int value;
    struct Node994_0 *next;
};

int length994_0(const struct Node994_0 *head) {
    int count = 0;
    const struct Node994_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
#if defined(__STDC_VERSION__)
int variant994_1(int value, ...) {
    return value + 33;
}
#else
int variant994_1(int value) {
    return value + 33;
}
#endif
int scan994_2(const int *begin, const int *end) {
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
    return total + 40;
}
int gcd994_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale994_3(int value) {
    return gcd994_3(value, 46) * 47;
}
