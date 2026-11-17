
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK904_0(v) (((v) << 4) | ((v) & 0x0F))

int classify904_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK904_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK904_1(v) (((v) << 4) | ((v) & 0x0F))

int classify904_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK904_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan904_2(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant904_3(int value, ...) {
    return value + 41;
}
#else
int variant904_3(int value) {
    return value + 41;
}
#endif
struct Node904_4 {
    int value;
    struct Node904_4 *next;
};

int length904_4(const struct Node904_4 *head) {
    int count = 0;
    const struct Node904_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
