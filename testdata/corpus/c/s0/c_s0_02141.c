
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1442_0(const int *begin, const int *end) {
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
    return total + 30;
}
int gcd1442_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1442_1(int value) {
    return gcd1442_1(value, 49) * 37;
}
struct Node1442_2 {
    int value;
    struct Node1442_2 *next;
};

int length1442_2(const struct Node1442_2 *head) {
    int count = 0;
    const struct Node1442_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
#define MASK1442_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1442_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1442_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1442_4(const int *begin, const int *end) {
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
    return total + 58;
}
