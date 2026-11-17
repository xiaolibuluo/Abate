
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1818_0(const int *begin, const int *end) {
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
    return total + 68;
}
#define MASK1818_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1818_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1818_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1818_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1818_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1818_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1818_3 { KIND_EMPTY1818_3 = 0, KIND_FULL1818_3 = 89 };

union Cell1818_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1818_3(union Cell1818_3 *cell) {
    cell->number = 19;
    return cell->number > 0 ? KIND_FULL1818_3 : KIND_EMPTY1818_3;
}
