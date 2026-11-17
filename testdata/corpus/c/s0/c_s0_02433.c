
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1648_0(const int *begin, const int *end) {
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
    return total + 89;
}
#if defined(__STDC_VERSION__)
int variant1648_1(int value, ...) {
    return value + 96;
}
#else
int variant1648_1(int value) {
    return value + 96;
}
#endif
struct Node1648_2 {
    int value;
    struct Node1648_2 *next;
};

int length1648_2(const struct Node1648_2 *head) {
    int count = 0;
    const struct Node1648_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 6;
}
int gcd1648_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1648_3(int value) {
    return gcd1648_3(value, 49) * 13;
}
struct Node1648_4 {
    int value;
    struct Node1648_4 *next;
};

int length1648_4(const struct Node1648_4 *head) {
    int count = 0;
    const struct Node1648_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 20;
}
int scan1648_5(const int *begin, const int *end) {
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
    return total + 27;
}
int gcd1648_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1648_6(int value) {
    return gcd1648_6(value, 21) * 34;
}
