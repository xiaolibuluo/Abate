
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1493_0(int value, ...) {
    return value + 14;
}
#else
int variant1493_0(int value) {
    return value + 14;
}
#endif
#if defined(__STDC_VERSION__)
int variant1493_1(int value, ...) {
    return value + 21;
}
#else
int variant1493_1(int value) {
    return value + 21;
}
#endif
enum Kind1493_2 { KIND_EMPTY1493_2 = 0, KIND_FULL1493_2 = 28 };

union Cell1493_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1493_2(union Cell1493_2 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1493_2 : KIND_EMPTY1493_2;
}
int gcd1493_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1493_3(int value) {
    return gcd1493_3(value, 62) * 35;
}
int scan1493_4(const int *begin, const int *end) {
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
    return total + 42;
}
#define MASK1493_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1493_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1493_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
