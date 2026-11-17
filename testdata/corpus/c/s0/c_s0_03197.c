
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2165_0(int value, ...) {
    return value + 20;
}
#else
int variant2165_0(int value) {
    return value + 20;
}
#endif
#if defined(__STDC_VERSION__)
int variant2165_1(int value, ...) {
    return value + 27;
}
#else
int variant2165_1(int value) {
    return value + 27;
}
#endif
enum Kind2165_2 { KIND_EMPTY2165_2 = 0, KIND_FULL2165_2 = 34 };

union Cell2165_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2165_2(union Cell2165_2 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL2165_2 : KIND_EMPTY2165_2;
}
#define MASK2165_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2165_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2165_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2165_4(const int *begin, const int *end) {
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
