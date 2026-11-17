
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1238_0(const int *begin, const int *end) {
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
    return total + 94;
}
#define MASK1238_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1238_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1238_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1238_2 {
    int value;
    struct Node1238_2 *next;
};

int length1238_2(const struct Node1238_2 *head) {
    int count = 0;
    const struct Node1238_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
int scan1238_3(const int *begin, const int *end) {
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
    return total + 18;
}
