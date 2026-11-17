
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2118_0(int value, ...) {
    return value + 88;
}
#else
int variant2118_0(int value) {
    return value + 88;
}
#endif
struct Node2118_1 {
    int value;
    struct Node2118_1 *next;
};

int length2118_1(const struct Node2118_1 *head) {
    int count = 0;
    const struct Node2118_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
#if defined(__STDC_VERSION__)
int variant2118_2(int value, ...) {
    return value + 5;
}
#else
int variant2118_2(int value) {
    return value + 5;
}
#endif
#define MASK2118_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2118_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2118_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2118_4(const int *begin, const int *end) {
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
    return total + 19;
}
