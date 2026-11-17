
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan364_0(const int *begin, const int *end) {
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
    return total + 81;
}
struct Node364_1 {
    int value;
    struct Node364_1 *next;
};

int length364_1(const struct Node364_1 *head) {
    int count = 0;
    const struct Node364_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
#define MASK364_2(v) (((v) << 4) | ((v) & 0x0F))

int classify364_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK364_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant364_3(int value, ...) {
    return value + 5;
}
#else
int variant364_3(int value) {
    return value + 5;
}
#endif
int scan364_4(const int *begin, const int *end) {
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
    return total + 12;
}
