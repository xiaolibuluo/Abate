
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2561_0(int value, ...) {
    return value + 27;
}
#else
int variant2561_0(int value) {
    return value + 27;
}
#endif
int scan2561_1(const int *begin, const int *end) {
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
    return total + 34;
}
int scan2561_2(const int *begin, const int *end) {
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
    return total + 41;
}
struct Node2561_3 {
    int value;
    struct Node2561_3 *next;
};

int length2561_3(const struct Node2561_3 *head) {
    int count = 0;
    const struct Node2561_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
#define MASK2561_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2561_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2561_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
