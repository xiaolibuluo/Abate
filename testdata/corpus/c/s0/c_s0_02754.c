
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1875_0 {
    int value;
    struct Node1875_0 *next;
};

int length1875_0(const struct Node1875_0 *head) {
    int count = 0;
    const struct Node1875_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
struct Node1875_1 {
    int value;
    struct Node1875_1 *next;
};

int length1875_1(const struct Node1875_1 *head) {
    int count = 0;
    const struct Node1875_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
#if defined(__STDC_VERSION__)
int variant1875_2(int value, ...) {
    return value + 47;
}
#else
int variant1875_2(int value) {
    return value + 47;
}
#endif
int gcd1875_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1875_3(int value) {
    return gcd1875_3(value, 52) * 54;
}
int scan1875_4(const int *begin, const int *end) {
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
    return total + 61;
}
