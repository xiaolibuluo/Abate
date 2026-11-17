
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2308_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2308_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2308_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2308_1 { KIND_EMPTY2308_1 = 0, KIND_FULL2308_1 = 43 };

union Cell2308_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2308_1(union Cell2308_1 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL2308_1 : KIND_EMPTY2308_1;
}
int scan2308_2(const int *begin, const int *end) {
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
    return total + 50;
}
#define MASK2308_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2308_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2308_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
