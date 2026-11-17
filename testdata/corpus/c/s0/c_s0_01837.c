
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1234_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1234_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1234_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1234_1(const int *begin, const int *end) {
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
enum Kind1234_2 { KIND_EMPTY1234_2 = 0, KIND_FULL1234_2 = 56 };

union Cell1234_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1234_2(union Cell1234_2 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1234_2 : KIND_EMPTY1234_2;
}
