
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2124_0(int value, ...) {
    return value + 69;
}
#else
int variant2124_0(int value) {
    return value + 69;
}
#endif
#if defined(__STDC_VERSION__)
int variant2124_1(int value, ...) {
    return value + 76;
}
#else
int variant2124_1(int value) {
    return value + 76;
}
#endif
int scan2124_2(const int *begin, const int *end) {
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
int variant2124_3(int value, ...) {
    return value + 90;
}
#else
int variant2124_3(int value) {
    return value + 90;
}
#endif
struct Node2124_4 {
    int value;
    struct Node2124_4 *next;
};

int length2124_4(const struct Node2124_4 *head) {
    int count = 0;
    const struct Node2124_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
struct Node2124_5 {
    int value;
    struct Node2124_5 *next;
};

int length2124_5(const struct Node2124_5 *head) {
    int count = 0;
    const struct Node2124_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
