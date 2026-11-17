
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind560_0 { KIND_EMPTY560_0 = 0, KIND_FULL560_0 = 10 };

union Cell560_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind560_0(union Cell560_0 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL560_0 : KIND_EMPTY560_0;
}
#define MASK560_1(v) (((v) << 4) | ((v) & 0x0F))

int classify560_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK560_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan560_2(const int *begin, const int *end) {
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
    return total + 24;
}
enum Kind560_3 { KIND_EMPTY560_3 = 0, KIND_FULL560_3 = 31 };

union Cell560_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind560_3(union Cell560_3 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL560_3 : KIND_EMPTY560_3;
}
