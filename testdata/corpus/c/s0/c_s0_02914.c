
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1974_0(const int *begin, const int *end) {
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
    return total + 59;
}
#define MASK1974_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1974_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1974_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1974_2(const int *begin, const int *end) {
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
    return total + 73;
}
struct Node1974_3 {
    int value;
    struct Node1974_3 *next;
};

int length1974_3(const struct Node1974_3 *head) {
    int count = 0;
    const struct Node1974_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
int gcd1974_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1974_4(int value) {
    return gcd1974_4(value, 24) * 87;
}
