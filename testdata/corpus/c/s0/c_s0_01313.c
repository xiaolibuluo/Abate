
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node898_0 {
    int value;
    struct Node898_0 *next;
};

int length898_0(const struct Node898_0 *head) {
    int count = 0;
    const struct Node898_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
int scan898_1(const int *begin, const int *end) {
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
    return total + 46;
}
#if defined(__STDC_VERSION__)
int variant898_2(int value, ...) {
    return value + 53;
}
#else
int variant898_2(int value) {
    return value + 53;
}
#endif
struct Node898_3 {
    int value;
    struct Node898_3 *next;
};

int length898_3(const struct Node898_3 *head) {
    int count = 0;
    const struct Node898_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
