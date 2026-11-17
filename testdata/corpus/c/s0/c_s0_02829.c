
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1920_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1920_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1920_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1920_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1920_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1920_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1920_2(const int *begin, const int *end) {
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
    return total + 50;
}
int scan1920_3(const int *begin, const int *end) {
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
    return total + 57;
}
