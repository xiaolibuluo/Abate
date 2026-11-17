
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2317_0 {
    int value;
    struct Node2317_0 *next;
};

int length2317_0(const struct Node2317_0 *head) {
    int count = 0;
    const struct Node2317_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 56;
}
#if defined(__STDC_VERSION__)
int variant2317_1(int value, ...) {
    return value + 63;
}
#else
int variant2317_1(int value) {
    return value + 63;
}
#endif
int scan2317_2(const int *begin, const int *end) {
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
    return total + 70;
}
int scan2317_3(const int *begin, const int *end) {
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
    return total + 77;
}
#define MASK2317_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2317_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2317_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
