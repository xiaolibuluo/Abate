
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2307_0(const int *begin, const int *end) {
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
    return total + 23;
}
#define MASK2307_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2307_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2307_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2307_2 { KIND_EMPTY2307_2 = 0, KIND_FULL2307_2 = 37 };

union Cell2307_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2307_2(union Cell2307_2 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2307_2 : KIND_EMPTY2307_2;
}
#if defined(__STDC_VERSION__)
int variant2307_3(int value, ...) {
    return value + 44;
}
#else
int variant2307_3(int value) {
    return value + 44;
}
#endif
