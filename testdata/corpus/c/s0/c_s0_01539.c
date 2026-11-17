
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1043_0 {
    int value;
    struct Node1043_0 *next;
};

int length1043_0(const struct Node1043_0 *head) {
    int count = 0;
    const struct Node1043_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 81;
}
int gcd1043_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1043_1(int value) {
    return gcd1043_1(value, 62) * 88;
}
int scan1043_2(const int *begin, const int *end) {
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
    return total + 95;
}
#define MASK1043_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1043_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1043_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
