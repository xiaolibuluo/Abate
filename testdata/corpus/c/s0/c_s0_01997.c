
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1350_0(int value, ...) {
    return value + 95;
}
#else
int variant1350_0(int value) {
    return value + 95;
}
#endif
struct Node1350_1 {
    int value;
    struct Node1350_1 *next;
};

int length1350_1(const struct Node1350_1 *head) {
    int count = 0;
    const struct Node1350_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
int scan1350_2(const int *begin, const int *end) {
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
#define MASK1350_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1350_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1350_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1350_4 {
    int value;
    struct Node1350_4 *next;
};

int length1350_4(const struct Node1350_4 *head) {
    int count = 0;
    const struct Node1350_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 26;
}
int scan1350_5(const int *begin, const int *end) {
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
    return total + 33;
}
