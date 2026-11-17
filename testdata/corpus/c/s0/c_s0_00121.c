
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan83_0(const int *begin, const int *end) {
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
    return total + 17;
}
#if defined(__STDC_VERSION__)
int variant83_1(int value, ...) {
    return value + 24;
}
#else
int variant83_1(int value) {
    return value + 24;
}
#endif
#define MASK83_2(v) (((v) << 4) | ((v) & 0x0F))

int classify83_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK83_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd83_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale83_3(int value) {
    return gcd83_3(value, 13) * 38;
}
enum Kind83_4 { KIND_EMPTY83_4 = 0, KIND_FULL83_4 = 45 };

union Cell83_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind83_4(union Cell83_4 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL83_4 : KIND_EMPTY83_4;
}
#define MASK83_5(v) (((v) << 4) | ((v) & 0x0F))

int classify83_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK83_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
