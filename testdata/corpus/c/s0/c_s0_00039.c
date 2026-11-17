
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant25_0(int value, ...) {
    return value + 39;
}
#else
int variant25_0(int value) {
    return value + 39;
}
#endif
int scan25_1(const int *begin, const int *end) {
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
#define MASK25_2(v) (((v) << 4) | ((v) & 0x0F))

int classify25_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK25_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node25_3 {
    int value;
    struct Node25_3 *next;
};

int length25_3(const struct Node25_3 *head) {
    int count = 0;
    const struct Node25_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
#if defined(__STDC_VERSION__)
int variant25_4(int value, ...) {
    return value + 67;
}
#else
int variant25_4(int value) {
    return value + 67;
}
#endif
