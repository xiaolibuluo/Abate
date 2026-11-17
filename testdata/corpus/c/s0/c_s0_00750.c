
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK517_0(v) (((v) << 4) | ((v) & 0x0F))

int classify517_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK517_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK517_1(v) (((v) << 4) | ((v) & 0x0F))

int classify517_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK517_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK517_2(v) (((v) << 4) | ((v) & 0x0F))

int classify517_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK517_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan517_3(const int *begin, const int *end) {
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
    return total + 54;
}
int gcd517_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale517_4(int value) {
    return gcd517_4(value, 12) * 61;
}
int scan517_5(const int *begin, const int *end) {
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
    return total + 68;
}
