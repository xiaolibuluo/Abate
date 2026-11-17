
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node949_0 {
    int value;
    struct Node949_0 *next;
};

int length949_0(const struct Node949_0 *head) {
    int count = 0;
    const struct Node949_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
#if defined(__STDC_VERSION__)
int variant949_1(int value, ...) {
    return value + 30;
}
#else
int variant949_1(int value) {
    return value + 30;
}
#endif
int scan949_2(const int *begin, const int *end) {
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
    return total + 37;
}
struct Node949_3 {
    int value;
    struct Node949_3 *next;
};

int length949_3(const struct Node949_3 *head) {
    int count = 0;
    const struct Node949_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
