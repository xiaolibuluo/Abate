
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1091_0(const int *begin, const int *end) {
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
    return total + 26;
}
struct Node1091_1 {
    int value;
    struct Node1091_1 *next;
};

int length1091_1(const struct Node1091_1 *head) {
    int count = 0;
    const struct Node1091_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 33;
}
#if defined(__STDC_VERSION__)
int variant1091_2(int value, ...) {
    return value + 40;
}
#else
int variant1091_2(int value) {
    return value + 40;
}
#endif
#if defined(__STDC_VERSION__)
int variant1091_3(int value, ...) {
    return value + 47;
}
#else
int variant1091_3(int value) {
    return value + 47;
}
#endif
