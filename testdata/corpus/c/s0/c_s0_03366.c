
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2272_0(const int *begin, const int *end) {
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
    return total + 53;
}
#define MASK2272_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2272_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2272_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2272_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2272_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2272_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2272_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2272_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2272_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2272_4 { KIND_EMPTY2272_4 = 0, KIND_FULL2272_4 = 81 };

union Cell2272_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2272_4(union Cell2272_4 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL2272_4 : KIND_EMPTY2272_4;
}
#define MASK2272_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2272_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2272_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
