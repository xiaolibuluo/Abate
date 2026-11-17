
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan275_0(const int *begin, const int *end) {
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
    return total + 88;
}
int scan275_1(const int *begin, const int *end) {
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
#define MASK275_2(v) (((v) << 4) | ((v) & 0x0F))

int classify275_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK275_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant275_3(int value, ...) {
    return value + 12;
}
#else
int variant275_3(int value) {
    return value + 12;
}
#endif
#define MASK275_4(v) (((v) << 4) | ((v) & 0x0F))

int classify275_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK275_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node275_5 {
    int value;
    struct Node275_5 *next;
};

int length275_5(const struct Node275_5 *head) {
    int count = 0;
    const struct Node275_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
