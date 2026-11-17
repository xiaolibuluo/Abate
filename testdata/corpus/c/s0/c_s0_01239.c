
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node843_0 {
    int value;
    struct Node843_0 *next;
};

int length843_0(const struct Node843_0 *head) {
    int count = 0;
    const struct Node843_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
int scan843_1(const int *begin, const int *end) {
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
    return total + 10;
}
#define MASK843_2(v) (((v) << 4) | ((v) & 0x0F))

int classify843_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK843_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
