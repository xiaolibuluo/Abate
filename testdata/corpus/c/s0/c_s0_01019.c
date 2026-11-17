
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd699_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale699_0(int value) {
    return gcd699_0(value, 22) * 71;
}
#define MASK699_1(v) (((v) << 4) | ((v) & 0x0F))

int classify699_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK699_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node699_2 {
    int value;
    struct Node699_2 *next;
};

int length699_2(const struct Node699_2 *head) {
    int count = 0;
    const struct Node699_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 85;
}
int scan699_3(const int *begin, const int *end) {
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
