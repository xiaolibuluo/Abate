
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd474_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale474_0(int value) {
    return gcd474_0(value, 33) * 56;
}
enum Kind474_1 { KIND_EMPTY474_1 = 0, KIND_FULL474_1 = 63 };

union Cell474_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind474_1(union Cell474_1 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL474_1 : KIND_EMPTY474_1;
}
#if defined(__STDC_VERSION__)
int variant474_2(int value, ...) {
    return value + 70;
}
#else
int variant474_2(int value) {
    return value + 70;
}
#endif
int scan474_3(const int *begin, const int *end) {
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
    return total + 77;
}
#define MASK474_4(v) (((v) << 4) | ((v) & 0x0F))

int classify474_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK474_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK474_5(v) (((v) << 4) | ((v) & 0x0F))

int classify474_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK474_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd474_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale474_6(int value) {
    return gcd474_6(value, 38) * 98;
}
