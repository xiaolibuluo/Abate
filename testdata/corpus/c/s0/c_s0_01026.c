
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK703_0(v) (((v) << 4) | ((v) & 0x0F))

int classify703_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK703_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan703_1(const int *begin, const int *end) {
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
    return total + 33;
}
#define MASK703_2(v) (((v) << 4) | ((v) & 0x0F))

int classify703_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK703_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan703_3(const int *begin, const int *end) {
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
    return total + 47;
}
