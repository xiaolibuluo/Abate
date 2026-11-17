
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2575_0 {
    int value;
    struct Node2575_0 *next;
};

int length2575_0(const struct Node2575_0 *head) {
    int count = 0;
    const struct Node2575_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#define MASK2575_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2575_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2575_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2575_2(const int *begin, const int *end) {
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
    return total + 29;
}
#define MASK2575_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2575_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2575_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2575_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2575_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2575_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2575_5 {
    int value;
    struct Node2575_5 *next;
};

int length2575_5(const struct Node2575_5 *head) {
    int count = 0;
    const struct Node2575_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
int scan2575_6(const int *begin, const int *end) {
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
    return total + 57;
}
