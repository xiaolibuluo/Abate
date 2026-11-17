
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan893_0(const int *begin, const int *end) {
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
    return total + 71;
}
#define MASK893_1(v) (((v) << 4) | ((v) & 0x0F))

int classify893_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK893_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd893_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale893_2(int value) {
    return gcd893_2(value, 60) * 85;
}
