
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK469_0(v) (((v) << 4) | ((v) & 0x0F))

int classify469_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK469_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node469_1 {
    int value;
    struct Node469_1 *next;
};

int length469_1(const struct Node469_1 *head) {
    int count = 0;
    const struct Node469_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 95;
}
int scan469_2(const int *begin, const int *end) {
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
    return total + 5;
}
int scan469_3(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant469_4(int value, ...) {
    return value + 19;
}
#else
int variant469_4(int value) {
    return value + 19;
}
#endif
