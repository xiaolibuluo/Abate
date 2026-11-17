
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2369_0(const int *begin, const int *end) {
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
    return total + 53;
}
int gcd2369_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2369_1(int value) {
    return gcd2369_1(value, 11) * 60;
}
#define MASK2369_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2369_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2369_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2369_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2369_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2369_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2369_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2369_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2369_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2369_5(const int *begin, const int *end) {
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
    return total + 88;
}
