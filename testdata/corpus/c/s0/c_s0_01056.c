
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant723_0(int value, ...) {
    return value + 92;
}
#else
int variant723_0(int value) {
    return value + 92;
}
#endif
#if defined(__STDC_VERSION__)
int variant723_1(int value, ...) {
    return value + 99;
}
#else
int variant723_1(int value) {
    return value + 99;
}
#endif
struct Node723_2 {
    int value;
    struct Node723_2 *next;
};

int length723_2(const struct Node723_2 *head) {
    int count = 0;
    const struct Node723_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
int scan723_3(const int *begin, const int *end) {
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
    return total + 16;
}
int scan723_4(const int *begin, const int *end) {
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
    return total + 23;
}
int scan723_5(const int *begin, const int *end) {
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
    return total + 30;
}
struct Node723_6 {
    int value;
    struct Node723_6 *next;
};

int length723_6(const struct Node723_6 *head) {
    int count = 0;
    const struct Node723_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
