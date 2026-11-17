
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan312_0(const int *begin, const int *end) {
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
#define MASK312_1(v) (((v) << 4) | ((v) & 0x0F))

int classify312_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK312_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind312_2 { KIND_EMPTY312_2 = 0, KIND_FULL312_2 = 98 };

union Cell312_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind312_2(union Cell312_2 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL312_2 : KIND_EMPTY312_2;
}
#if defined(__STDC_VERSION__)
int variant312_3(int value, ...) {
    return value + 8;
}
#else
int variant312_3(int value) {
    return value + 8;
}
#endif
