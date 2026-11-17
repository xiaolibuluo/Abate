
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK561_0(v) (((v) << 4) | ((v) & 0x0F))

int classify561_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK561_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node561_1 {
    int value;
    struct Node561_1 *next;
};

int length561_1(const struct Node561_1 *head) {
    int count = 0;
    const struct Node561_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
#define MASK561_2(v) (((v) << 4) | ((v) & 0x0F))

int classify561_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK561_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant561_3(int value, ...) {
    return value + 44;
}
#else
int variant561_3(int value) {
    return value + 44;
}
#endif
int scan561_4(const int *begin, const int *end) {
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
    return total + 51;
}
