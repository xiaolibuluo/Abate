
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node758_0 {
    int value;
    struct Node758_0 *next;
};

int length758_0(const struct Node758_0 *head) {
    int count = 0;
    const struct Node758_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
#define MASK758_1(v) (((v) << 4) | ((v) & 0x0F))

int classify758_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK758_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant758_2(int value, ...) {
    return value + 76;
}
#else
int variant758_2(int value) {
    return value + 76;
}
#endif
int scan758_3(const int *begin, const int *end) {
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
    return total + 83;
}
#if defined(__STDC_VERSION__)
int variant758_4(int value, ...) {
    return value + 90;
}
#else
int variant758_4(int value) {
    return value + 90;
}
#endif
struct Node758_5 {
    int value;
    struct Node758_5 *next;
};

int length758_5(const struct Node758_5 *head) {
    int count = 0;
    const struct Node758_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
