
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2363_0(int value, ...) {
    return value + 72;
}
#else
int variant2363_0(int value) {
    return value + 72;
}
#endif
struct Node2363_1 {
    int value;
    struct Node2363_1 *next;
};

int length2363_1(const struct Node2363_1 *head) {
    int count = 0;
    const struct Node2363_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
int scan2363_2(const int *begin, const int *end) {
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
    return total + 86;
}
#define MASK2363_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2363_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2363_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2363_4(int value, ...) {
    return value + 100;
}
#else
int variant2363_4(int value) {
    return value + 100;
}
#endif
