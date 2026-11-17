
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd924_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale924_0(int value) {
    return gcd924_0(value, 11) * 86;
}
#if defined(__STDC_VERSION__)
int variant924_1(int value, ...) {
    return value + 93;
}
#else
int variant924_1(int value) {
    return value + 93;
}
#endif
int scan924_2(const int *begin, const int *end) {
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
    return total + 100;
}
enum Kind924_3 { KIND_EMPTY924_3 = 0, KIND_FULL924_3 = 10 };

union Cell924_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind924_3(union Cell924_3 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL924_3 : KIND_EMPTY924_3;
}
int gcd924_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale924_4(int value) {
    return gcd924_4(value, 55) * 17;
}
#define MASK924_5(v) (((v) << 4) | ((v) & 0x0F))

int classify924_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK924_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind924_6 { KIND_EMPTY924_6 = 0, KIND_FULL924_6 = 31 };

union Cell924_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind924_6(union Cell924_6 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL924_6 : KIND_EMPTY924_6;
}
enum Kind924_7 { KIND_EMPTY924_7 = 0, KIND_FULL924_7 = 38 };

union Cell924_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind924_7(union Cell924_7 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL924_7 : KIND_EMPTY924_7;
}
