
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan862_0(const int *begin, const int *end) {
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
    return total + 56;
}
int scan862_1(const int *begin, const int *end) {
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
    return total + 63;
}
#if defined(__STDC_VERSION__)
int variant862_2(int value, ...) {
    return value + 70;
}
#else
int variant862_2(int value) {
    return value + 70;
}
#endif
int gcd862_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale862_3(int value) {
    return gcd862_3(value, 37) * 77;
}
#if defined(__STDC_VERSION__)
int variant862_4(int value, ...) {
    return value + 84;
}
#else
int variant862_4(int value) {
    return value + 84;
}
#endif
#define MASK862_5(v) (((v) << 4) | ((v) & 0x0F))

int classify862_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK862_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind862_6 { KIND_EMPTY862_6 = 0, KIND_FULL862_6 = 98 };

union Cell862_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind862_6(union Cell862_6 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL862_6 : KIND_EMPTY862_6;
}
