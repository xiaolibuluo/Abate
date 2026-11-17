
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node856_0 {
    int value;
    struct Node856_0 *next;
};

int length856_0(const struct Node856_0 *head) {
    int count = 0;
    const struct Node856_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
struct Node856_1 {
    int value;
    struct Node856_1 *next;
};

int length856_1(const struct Node856_1 *head) {
    int count = 0;
    const struct Node856_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 82;
}
#if defined(__STDC_VERSION__)
int variant856_2(int value, ...) {
    return value + 89;
}
#else
int variant856_2(int value) {
    return value + 89;
}
#endif
#if defined(__STDC_VERSION__)
int variant856_3(int value, ...) {
    return value + 96;
}
#else
int variant856_3(int value) {
    return value + 96;
}
#endif
int gcd856_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale856_4(int value) {
    return gcd856_4(value, 67) * 6;
}
int scan856_5(const int *begin, const int *end) {
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
    return total + 13;
}
