
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2496_0(int value, ...) {
    return value + 55;
}
#else
int variant2496_0(int value) {
    return value + 55;
}
#endif
int scan2496_1(const int *begin, const int *end) {
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
    return total + 62;
}
struct Node2496_2 {
    int value;
    struct Node2496_2 *next;
};

int length2496_2(const struct Node2496_2 *head) {
    int count = 0;
    const struct Node2496_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
