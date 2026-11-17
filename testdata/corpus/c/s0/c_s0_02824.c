
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1918_0 {
    int value;
    struct Node1918_0 *next;
};

int length1918_0(const struct Node1918_0 *head) {
    int count = 0;
    const struct Node1918_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
int scan1918_1(const int *begin, const int *end) {
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
    return total + 17;
}
#if defined(__STDC_VERSION__)
int variant1918_2(int value, ...) {
    return value + 24;
}
#else
int variant1918_2(int value) {
    return value + 24;
}
#endif
int scan1918_3(const int *begin, const int *end) {
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
    return total + 31;
}
