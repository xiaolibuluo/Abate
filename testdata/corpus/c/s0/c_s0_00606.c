
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant414_0(int value, ...) {
    return value + 52;
}
#else
int variant414_0(int value) {
    return value + 52;
}
#endif
#define MASK414_1(v) (((v) << 4) | ((v) & 0x0F))

int classify414_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK414_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant414_2(int value, ...) {
    return value + 66;
}
#else
int variant414_2(int value) {
    return value + 66;
}
#endif
int scan414_3(const int *begin, const int *end) {
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
    return total + 73;
}
#if defined(__STDC_VERSION__)
int variant414_4(int value, ...) {
    return value + 80;
}
#else
int variant414_4(int value) {
    return value + 80;
}
#endif
struct Node414_5 {
    int value;
    struct Node414_5 *next;
};

int length414_5(const struct Node414_5 *head) {
    int count = 0;
    const struct Node414_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 87;
}
