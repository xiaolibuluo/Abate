
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant468_0(int value, ...) {
    return value + 75;
}
#else
int variant468_0(int value) {
    return value + 75;
}
#endif
#define MASK468_1(v) (((v) << 4) | ((v) & 0x0F))

int classify468_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK468_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan468_2(const int *begin, const int *end) {
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
    return total + 89;
}
int scan468_3(const int *begin, const int *end) {
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
    return total + 96;
}
enum Kind468_4 { KIND_EMPTY468_4 = 0, KIND_FULL468_4 = 6 };

union Cell468_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind468_4(union Cell468_4 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL468_4 : KIND_EMPTY468_4;
}
