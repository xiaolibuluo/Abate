
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1927_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1927_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1927_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1927_1 {
    int value;
    struct Node1927_1 *next;
};

int length1927_1(const struct Node1927_1 *head) {
    int count = 0;
    const struct Node1927_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
int scan1927_2(const int *begin, const int *end) {
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
    return total + 44;
}
int scan1927_3(const int *begin, const int *end) {
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
struct Node1927_4 {
    int value;
    struct Node1927_4 *next;
};

int length1927_4(const struct Node1927_4 *head) {
    int count = 0;
    const struct Node1927_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
