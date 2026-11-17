
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2461_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2461_0(int value) {
    return gcd2461_0(value, 34) * 85;
}
int scan2461_1(const int *begin, const int *end) {
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
    return total + 92;
}
#define MASK2461_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2461_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2461_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2461_3(const int *begin, const int *end) {
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
    return total + 9;
}
struct Node2461_4 {
    int value;
    struct Node2461_4 *next;
};

int length2461_4(const struct Node2461_4 *head) {
    int count = 0;
    const struct Node2461_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 16;
}
