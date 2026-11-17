
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1393_0 {
    int value;
    struct Node1393_0 *next;
};

int length1393_0(const struct Node1393_0 *head) {
    int count = 0;
    const struct Node1393_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
struct Node1393_1 {
    int value;
    struct Node1393_1 *next;
};

int length1393_1(const struct Node1393_1 *head) {
    int count = 0;
    const struct Node1393_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
int gcd1393_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1393_2(int value) {
    return gcd1393_2(value, 22) * 86;
}
#if defined(__STDC_VERSION__)
int variant1393_3(int value, ...) {
    return value + 93;
}
#else
int variant1393_3(int value) {
    return value + 93;
}
#endif
int scan1393_4(const int *begin, const int *end) {
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
    return total + 100;
}
