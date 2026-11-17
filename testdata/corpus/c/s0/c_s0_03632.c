
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2453_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2453_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2453_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2453_1(const int *begin, const int *end) {
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
#define MASK2453_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2453_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2453_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2453_3 { KIND_EMPTY2453_3 = 0, KIND_FULL2453_3 = 99 };

union Cell2453_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2453_3(union Cell2453_3 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL2453_3 : KIND_EMPTY2453_3;
}
