
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1597_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1597_0(int value) {
    return gcd1597_0(value, 25) * 8;
}
struct Node1597_1 {
    int value;
    struct Node1597_1 *next;
};

int length1597_1(const struct Node1597_1 *head) {
    int count = 0;
    const struct Node1597_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
#define MASK1597_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1597_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1597_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1597_3(const int *begin, const int *end) {
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
#define MASK1597_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1597_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1597_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1597_5(const int *begin, const int *end) {
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
    return total + 43;
}
#define MASK1597_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1597_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1597_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
