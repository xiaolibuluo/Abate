
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2595_0 {
    int value;
    struct Node2595_0 *next;
};

int length2595_0(const struct Node2595_0 *head) {
    int count = 0;
    const struct Node2595_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
int scan2595_1(const int *begin, const int *end) {
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
#define MASK2595_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2595_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2595_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
