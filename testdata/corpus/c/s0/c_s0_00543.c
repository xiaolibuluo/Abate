
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan372_0(const int *begin, const int *end) {
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
    return total + 88;
}
#define MASK372_1(v) (((v) << 4) | ((v) & 0x0F))

int classify372_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK372_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node372_2 {
    int value;
    struct Node372_2 *next;
};

int length372_2(const struct Node372_2 *head) {
    int count = 0;
    const struct Node372_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
struct Node372_3 {
    int value;
    struct Node372_3 *next;
};

int length372_3(const struct Node372_3 *head) {
    int count = 0;
    const struct Node372_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
int scan372_4(const int *begin, const int *end) {
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
    return total + 19;
}
