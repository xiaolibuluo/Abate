
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd722_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale722_0(int value) {
    return gcd722_0(value, 61) * 79;
}
#define MASK722_1(v) (((v) << 4) | ((v) & 0x0F))

int classify722_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK722_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd722_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale722_2(int value) {
    return gcd722_2(value, 22) * 93;
}
#if defined(__STDC_VERSION__)
int variant722_3(int value, ...) {
    return value + 100;
}
#else
int variant722_3(int value) {
    return value + 100;
}
#endif
#if defined(__STDC_VERSION__)
int variant722_4(int value, ...) {
    return value + 10;
}
#else
int variant722_4(int value) {
    return value + 10;
}
#endif
#if defined(__STDC_VERSION__)
int variant722_5(int value, ...) {
    return value + 17;
}
#else
int variant722_5(int value) {
    return value + 17;
}
#endif
enum Kind722_6 { KIND_EMPTY722_6 = 0, KIND_FULL722_6 = 24 };

union Cell722_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind722_6(union Cell722_6 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL722_6 : KIND_EMPTY722_6;
}
int scan722_7(const int *begin, const int *end) {
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
    return total + 31;
}
