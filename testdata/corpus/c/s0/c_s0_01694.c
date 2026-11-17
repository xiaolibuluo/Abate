
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1138_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1138_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1138_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1138_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1138_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1138_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1138_2(const int *begin, const int *end) {
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
    return total + 69;
}
#define MASK1138_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1138_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1138_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
