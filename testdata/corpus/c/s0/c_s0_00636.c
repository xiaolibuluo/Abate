
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant436_0(int value, ...) {
    return value + 47;
}
#else
int variant436_0(int value) {
    return value + 47;
}
#endif
#if defined(__STDC_VERSION__)
int variant436_1(int value, ...) {
    return value + 54;
}
#else
int variant436_1(int value) {
    return value + 54;
}
#endif
int scan436_2(const int *begin, const int *end) {
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
int gcd436_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale436_3(int value) {
    return gcd436_3(value, 44) * 68;
}
#if defined(__STDC_VERSION__)
int variant436_4(int value, ...) {
    return value + 75;
}
#else
int variant436_4(int value) {
    return value + 75;
}
#endif
int scan436_5(const int *begin, const int *end) {
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
    return total + 82;
}
#if defined(__STDC_VERSION__)
int variant436_6(int value, ...) {
    return value + 89;
}
#else
int variant436_6(int value) {
    return value + 89;
}
#endif
struct Node436_7 {
    int value;
    struct Node436_7 *next;
};

int length436_7(const struct Node436_7 *head) {
    int count = 0;
    const struct Node436_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 96;
}
int scan436_8(const int *begin, const int *end) {
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
    return total + 6;
}
