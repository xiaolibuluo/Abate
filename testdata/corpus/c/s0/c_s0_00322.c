
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant222_0(int value, ...) {
    return value + 78;
}
#else
int variant222_0(int value) {
    return value + 78;
}
#endif
int scan222_1(const int *begin, const int *end) {
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
    return total + 85;
}
#if defined(__STDC_VERSION__)
int variant222_2(int value, ...) {
    return value + 92;
}
#else
int variant222_2(int value) {
    return value + 92;
}
#endif
#define MASK222_3(v) (((v) << 4) | ((v) & 0x0F))

int classify222_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK222_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind222_4 { KIND_EMPTY222_4 = 0, KIND_FULL222_4 = 9 };

union Cell222_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind222_4(union Cell222_4 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL222_4 : KIND_EMPTY222_4;
}
enum Kind222_5 { KIND_EMPTY222_5 = 0, KIND_FULL222_5 = 16 };

union Cell222_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind222_5(union Cell222_5 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL222_5 : KIND_EMPTY222_5;
}
