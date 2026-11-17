
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant409_0(int value, ...) {
    return value + 84;
}
#else
int variant409_0(int value) {
    return value + 84;
}
#endif
int scan409_1(const int *begin, const int *end) {
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
    return total + 91;
}
#define MASK409_2(v) (((v) << 4) | ((v) & 0x0F))

int classify409_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK409_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node409_3 {
    int value;
    struct Node409_3 *next;
};

int length409_3(const struct Node409_3 *head) {
    int count = 0;
    const struct Node409_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 8;
}
