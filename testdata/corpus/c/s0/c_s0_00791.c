
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan545_0(const int *begin, const int *end) {
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
    return total + 9;
}
#if defined(__STDC_VERSION__)
int variant545_1(int value, ...) {
    return value + 16;
}
#else
int variant545_1(int value) {
    return value + 16;
}
#endif
struct Node545_2 {
    int value;
    struct Node545_2 *next;
};

int length545_2(const struct Node545_2 *head) {
    int count = 0;
    const struct Node545_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
