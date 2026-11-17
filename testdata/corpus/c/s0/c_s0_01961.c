
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1324_0(const int *begin, const int *end) {
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
    return total + 48;
}
#define MASK1324_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1324_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1324_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
