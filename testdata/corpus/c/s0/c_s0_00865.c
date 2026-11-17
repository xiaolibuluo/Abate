
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node595_0 {
    int value;
    struct Node595_0 *next;
};

int length595_0(const struct Node595_0 *head) {
    int count = 0;
    const struct Node595_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
#if defined(__STDC_VERSION__)
int variant595_1(int value, ...) {
    return value + 84;
}
#else
int variant595_1(int value) {
    return value + 84;
}
#endif
int scan595_2(const int *begin, const int *end) {
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
    return total + 91;
}
