
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2565_0 {
    int value;
    struct Node2565_0 *next;
};

int length2565_0(const struct Node2565_0 *head) {
    int count = 0;
    const struct Node2565_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 79;
}
int scan2565_1(const int *begin, const int *end) {
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
    return total + 86;
}
#define MASK2565_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2565_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2565_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2565_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2565_3(int value) {
    return gcd2565_3(value, 63) * 100;
}
#define MASK2565_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2565_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2565_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2565_5 {
    int value;
    struct Node2565_5 *next;
};

int length2565_5(const struct Node2565_5 *head) {
    int count = 0;
    const struct Node2565_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
