
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan118_0(const int *begin, const int *end) {
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
    return total + 84;
}
#define MASK118_1(v) (((v) << 4) | ((v) & 0x0F))

int classify118_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK118_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd118_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale118_2(int value) {
    return gcd118_2(value, 64) * 98;
}
#define MASK118_3(v) (((v) << 4) | ((v) & 0x0F))

int classify118_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK118_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
