
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan894_0(const int *begin, const int *end) {
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
    return total + 84;
}
struct Node894_1 {
    int value;
    struct Node894_1 *next;
};

int length894_1(const struct Node894_1 *head) {
    int count = 0;
    const struct Node894_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
#define MASK894_2(v) (((v) << 4) | ((v) & 0x0F))

int classify894_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK894_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
