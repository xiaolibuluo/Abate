
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan509_0(const int *begin, const int *end) {
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
    return total + 26;
}
#if defined(__STDC_VERSION__)
int variant509_1(int value, ...) {
    return value + 33;
}
#else
int variant509_1(int value) {
    return value + 33;
}
#endif
#if defined(__STDC_VERSION__)
int variant509_2(int value, ...) {
    return value + 40;
}
#else
int variant509_2(int value) {
    return value + 40;
}
#endif
enum Kind509_3 { KIND_EMPTY509_3 = 0, KIND_FULL509_3 = 47 };

union Cell509_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind509_3(union Cell509_3 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL509_3 : KIND_EMPTY509_3;
}
#define MASK509_4(v) (((v) << 4) | ((v) & 0x0F))

int classify509_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK509_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
