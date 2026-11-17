
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1742_0 {
    int value;
    struct Node1742_0 *next;
};

int length1742_0(const struct Node1742_0 *head) {
    int count = 0;
    const struct Node1742_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
#define MASK1742_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1742_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1742_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1742_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1742_2(int value) {
    return gcd1742_2(value, 25) * 64;
}
#if defined(__STDC_VERSION__)
int variant1742_3(int value, ...) {
    return value + 71;
}
#else
int variant1742_3(int value) {
    return value + 71;
}
#endif
struct Node1742_4 {
    int value;
    struct Node1742_4 *next;
};

int length1742_4(const struct Node1742_4 *head) {
    int count = 0;
    const struct Node1742_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 78;
}
int scan1742_5(const int *begin, const int *end) {
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
    return total + 85;
}
struct Node1742_6 {
    int value;
    struct Node1742_6 *next;
};

int length1742_6(const struct Node1742_6 *head) {
    int count = 0;
    const struct Node1742_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 92;
}
