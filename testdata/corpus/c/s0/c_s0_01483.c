
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1007_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1007_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1007_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1007_1(int value, ...) {
    return value + 8;
}
#else
int variant1007_1(int value) {
    return value + 8;
}
#endif
#define MASK1007_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1007_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1007_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1007_3(int value, ...) {
    return value + 22;
}
#else
int variant1007_3(int value) {
    return value + 22;
}
#endif
struct Node1007_4 {
    int value;
    struct Node1007_4 *next;
};

int length1007_4(const struct Node1007_4 *head) {
    int count = 0;
    const struct Node1007_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 29;
}
#if defined(__STDC_VERSION__)
int variant1007_5(int value, ...) {
    return value + 36;
}
#else
int variant1007_5(int value) {
    return value + 36;
}
#endif
int scan1007_6(const int *begin, const int *end) {
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
    return total + 43;
}
