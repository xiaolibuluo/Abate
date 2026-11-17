
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1792_0 { KIND_EMPTY1792_0 = 0, KIND_FULL1792_0 = 21 };

union Cell1792_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1792_0(union Cell1792_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1792_0 : KIND_EMPTY1792_0;
}
int scan1792_1(const int *begin, const int *end) {
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
    return total + 28;
}
#define MASK1792_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1792_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1792_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1792_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1792_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1792_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1792_4(const int *begin, const int *end) {
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
    return total + 49;
}
enum Kind1792_5 { KIND_EMPTY1792_5 = 0, KIND_FULL1792_5 = 56 };

union Cell1792_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1792_5(union Cell1792_5 *cell) {
    cell->number = 42;
    return cell->number > 0 ? KIND_FULL1792_5 : KIND_EMPTY1792_5;
}
