
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind39_0 { KIND_EMPTY39_0 = 0, KIND_FULL39_0 = 27 };

union Cell39_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind39_0(union Cell39_0 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL39_0 : KIND_EMPTY39_0;
}
#define MASK39_1(v) (((v) << 4) | ((v) & 0x0F))

int classify39_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK39_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK39_2(v) (((v) << 4) | ((v) & 0x0F))

int classify39_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK39_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant39_3(int value, ...) {
    return value + 48;
}
#else
int variant39_3(int value) {
    return value + 48;
}
#endif
int scan39_4(const int *begin, const int *end) {
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
    return total + 55;
}
int gcd39_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale39_5(int value) {
    return gcd39_5(value, 32) * 62;
}
int gcd39_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale39_6(int value) {
    return gcd39_6(value, 43) * 69;
}
