
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1747_0(int value, ...) {
    return value + 18;
}
#else
int variant1747_0(int value) {
    return value + 18;
}
#endif
#if defined(__STDC_VERSION__)
int variant1747_1(int value, ...) {
    return value + 25;
}
#else
int variant1747_1(int value) {
    return value + 25;
}
#endif
struct Node1747_2 {
    int value;
    struct Node1747_2 *next;
};

int length1747_2(const struct Node1747_2 *head) {
    int count = 0;
    const struct Node1747_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
struct Node1747_3 {
    int value;
    struct Node1747_3 *next;
};

int length1747_3(const struct Node1747_3 *head) {
    int count = 0;
    const struct Node1747_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
int gcd1747_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1747_4(int value) {
    return gcd1747_4(value, 21) * 46;
}
int gcd1747_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1747_5(int value) {
    return gcd1747_5(value, 32) * 53;
}
struct Node1747_6 {
    int value;
    struct Node1747_6 *next;
};

int length1747_6(const struct Node1747_6 *head) {
    int count = 0;
    const struct Node1747_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
int scan1747_7(const int *begin, const int *end) {
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
    return total + 67;
}
