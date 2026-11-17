
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan952_0(const int *begin, const int *end) {
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
struct Node952_1 {
    int value;
    struct Node952_1 *next;
};

int length952_1(const struct Node952_1 *head) {
    int count = 0;
    const struct Node952_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
struct Node952_2 {
    int value;
    struct Node952_2 *next;
};

int length952_2(const struct Node952_2 *head) {
    int count = 0;
    const struct Node952_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 76;
}
#if defined(__STDC_VERSION__)
int variant952_3(int value, ...) {
    return value + 83;
}
#else
int variant952_3(int value) {
    return value + 83;
}
#endif
#if defined(__STDC_VERSION__)
int variant952_4(int value, ...) {
    return value + 90;
}
#else
int variant952_4(int value) {
    return value + 90;
}
#endif
