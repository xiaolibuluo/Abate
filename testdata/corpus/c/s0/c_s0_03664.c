
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2475_0 {
    int value;
    struct Node2475_0 *next;
};

int length2475_0(const struct Node2475_0 *head) {
    int count = 0;
    const struct Node2475_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
int scan2475_1(const int *begin, const int *end) {
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
    return total + 80;
}
int scan2475_2(const int *begin, const int *end) {
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
    return total + 87;
}
#if defined(__STDC_VERSION__)
int variant2475_3(int value, ...) {
    return value + 94;
}
#else
int variant2475_3(int value) {
    return value + 94;
}
#endif
#if defined(__STDC_VERSION__)
int variant2475_4(int value, ...) {
    return value + 101;
}
#else
int variant2475_4(int value) {
    return value + 101;
}
#endif
struct Node2475_5 {
    int value;
    struct Node2475_5 *next;
};

int length2475_5(const struct Node2475_5 *head) {
    int count = 0;
    const struct Node2475_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
#if defined(__STDC_VERSION__)
int variant2475_6(int value, ...) {
    return value + 18;
}
#else
int variant2475_6(int value) {
    return value + 18;
}
#endif
int scan2475_7(const int *begin, const int *end) {
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
    return total + 25;
}
