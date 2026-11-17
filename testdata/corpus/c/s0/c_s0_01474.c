
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1002_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1002_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1002_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1002_1 { KIND_EMPTY1002_1 = 0, KIND_FULL1002_1 = 40 };

union Cell1002_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1002_1(union Cell1002_1 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1002_1 : KIND_EMPTY1002_1;
}
#define MASK1002_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1002_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1002_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1002_3(const int *begin, const int *end) {
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
    return total + 54;
}
int scan1002_4(const int *begin, const int *end) {
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
    return total + 61;
}
