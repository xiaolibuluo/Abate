
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant410_0(int value, ...) {
    return value + 97;
}
#else
int variant410_0(int value) {
    return value + 97;
}
#endif
struct Node410_1 {
    int value;
    struct Node410_1 *next;
};

int length410_1(const struct Node410_1 *head) {
    int count = 0;
    const struct Node410_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
int scan410_2(const int *begin, const int *end) {
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
    return total + 14;
}
struct Node410_3 {
    int value;
    struct Node410_3 *next;
};

int length410_3(const struct Node410_3 *head) {
    int count = 0;
    const struct Node410_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
