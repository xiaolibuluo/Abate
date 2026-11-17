
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2021_0 {
    int value;
    struct Node2021_0 *next;
};

int length2021_0(const struct Node2021_0 *head) {
    int count = 0;
    const struct Node2021_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
int scan2021_1(const int *begin, const int *end) {
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
    return total + 95;
}
#if defined(__STDC_VERSION__)
int variant2021_2(int value, ...) {
    return value + 5;
}
#else
int variant2021_2(int value) {
    return value + 5;
}
#endif
#if defined(__STDC_VERSION__)
int variant2021_3(int value, ...) {
    return value + 12;
}
#else
int variant2021_3(int value) {
    return value + 12;
}
#endif
struct Node2021_4 {
    int value;
    struct Node2021_4 *next;
};

int length2021_4(const struct Node2021_4 *head) {
    int count = 0;
    const struct Node2021_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 19;
}
