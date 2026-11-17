
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2206_0(const int *begin, const int *end) {
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
    return total + 68;
}
#if defined(__STDC_VERSION__)
int variant2206_1(int value, ...) {
    return value + 75;
}
#else
int variant2206_1(int value) {
    return value + 75;
}
#endif
struct Node2206_2 {
    int value;
    struct Node2206_2 *next;
};

int length2206_2(const struct Node2206_2 *head) {
    int count = 0;
    const struct Node2206_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
