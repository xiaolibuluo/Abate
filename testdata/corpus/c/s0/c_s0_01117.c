
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node766_0 {
    int value;
    struct Node766_0 *next;
};

int length766_0(const struct Node766_0 *head) {
    int count = 0;
    const struct Node766_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
int gcd766_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale766_1(int value) {
    return gcd766_1(value, 14) * 76;
}
int scan766_2(const int *begin, const int *end) {
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
    return total + 83;
}
#if defined(__STDC_VERSION__)
int variant766_3(int value, ...) {
    return value + 90;
}
#else
int variant766_3(int value) {
    return value + 90;
}
#endif
int gcd766_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale766_4(int value) {
    return gcd766_4(value, 47) * 97;
}
int gcd766_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale766_5(int value) {
    return gcd766_5(value, 58) * 7;
}
#if defined(__STDC_VERSION__)
int variant766_6(int value, ...) {
    return value + 14;
}
#else
int variant766_6(int value) {
    return value + 14;
}
#endif
