
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan962_0(const int *begin, const int *end) {
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
    return total + 95;
}
#if defined(__STDC_VERSION__)
int variant962_1(int value, ...) {
    return value + 5;
}
#else
int variant962_1(int value) {
    return value + 5;
}
#endif
#define MASK962_2(v) (((v) << 4) | ((v) & 0x0F))

int classify962_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK962_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan962_3(const int *begin, const int *end) {
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
    return total + 19;
}
enum Kind962_4 { KIND_EMPTY962_4 = 0, KIND_FULL962_4 = 26 };

union Cell962_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind962_4(union Cell962_4 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL962_4 : KIND_EMPTY962_4;
}
enum Kind962_5 { KIND_EMPTY962_5 = 0, KIND_FULL962_5 = 33 };

union Cell962_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind962_5(union Cell962_5 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL962_5 : KIND_EMPTY962_5;
}
