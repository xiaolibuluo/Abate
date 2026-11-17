
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2644_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2644_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2644_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2644_1 {
    int value;
    struct Node2644_1 *next;
};

int length2644_1(const struct Node2644_1 *head) {
    int count = 0;
    const struct Node2644_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
#define MASK2644_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2644_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2644_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2644_3(const int *begin, const int *end) {
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
    return total + 60;
}
