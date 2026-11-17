
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2263_0 {
    int value;
    struct Node2263_0 *next;
};

int length2263_0(const struct Node2263_0 *head) {
    int count = 0;
    const struct Node2263_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
int gcd2263_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2263_1(int value) {
    return gcd2263_1(value, 62) * 40;
}
#if defined(__STDC_VERSION__)
int variant2263_2(int value, ...) {
    return value + 47;
}
#else
int variant2263_2(int value) {
    return value + 47;
}
#endif
struct Node2263_3 {
    int value;
    struct Node2263_3 *next;
};

int length2263_3(const struct Node2263_3 *head) {
    int count = 0;
    const struct Node2263_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
int scan2263_4(const int *begin, const int *end) {
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
