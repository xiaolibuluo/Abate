
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2279_0 {
    int value;
    struct Node2279_0 *next;
};

int length2279_0(const struct Node2279_0 *head) {
    int count = 0;
    const struct Node2279_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
struct Node2279_1 {
    int value;
    struct Node2279_1 *next;
};

int length2279_1(const struct Node2279_1 *head) {
    int count = 0;
    const struct Node2279_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
#if defined(__STDC_VERSION__)
int variant2279_2(int value, ...) {
    return value + 61;
}
#else
int variant2279_2(int value) {
    return value + 61;
}
#endif
#if defined(__STDC_VERSION__)
int variant2279_3(int value, ...) {
    return value + 68;
}
#else
int variant2279_3(int value) {
    return value + 68;
}
#endif
int scan2279_4(const int *begin, const int *end) {
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
    return total + 75;
}
int scan2279_5(const int *begin, const int *end) {
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
int scan2279_6(const int *begin, const int *end) {
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
