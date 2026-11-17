
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind663_0 { KIND_EMPTY663_0 = 0, KIND_FULL663_0 = 88 };

union Cell663_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind663_0(union Cell663_0 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL663_0 : KIND_EMPTY663_0;
}
int scan663_1(const int *begin, const int *end) {
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
#define MASK663_2(v) (((v) << 4) | ((v) & 0x0F))

int classify663_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK663_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant663_3(int value, ...) {
    return value + 12;
}
#else
int variant663_3(int value) {
    return value + 12;
}
#endif
enum Kind663_4 { KIND_EMPTY663_4 = 0, KIND_FULL663_4 = 19 };

union Cell663_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind663_4(union Cell663_4 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL663_4 : KIND_EMPTY663_4;
}
#if defined(__STDC_VERSION__)
int variant663_5(int value, ...) {
    return value + 26;
}
#else
int variant663_5(int value) {
    return value + 26;
}
#endif
