
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK553_0(v) (((v) << 4) | ((v) & 0x0F))

int classify553_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK553_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan553_1(const int *begin, const int *end) {
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
    return total + 23;
}
int gcd553_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale553_2(int value) {
    return gcd553_2(value, 59) * 30;
}
struct Node553_3 {
    int value;
    struct Node553_3 *next;
};

int length553_3(const struct Node553_3 *head) {
    int count = 0;
    const struct Node553_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
