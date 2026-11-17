
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1106_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1106_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1106_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1106_1(const int *begin, const int *end) {
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
    return total + 34;
}
#define MASK1106_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1106_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1106_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1106_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1106_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1106_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1106_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1106_4(int value) {
    return gcd1106_4(value, 48) * 55;
}
int scan1106_5(const int *begin, const int *end) {
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
    return total + 62;
}
