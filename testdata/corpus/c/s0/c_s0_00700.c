
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant480_0(int value, ...) {
    return value + 37;
}
#else
int variant480_0(int value) {
    return value + 37;
}
#endif
#define MASK480_1(v) (((v) << 4) | ((v) & 0x0F))

int classify480_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK480_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK480_2(v) (((v) << 4) | ((v) & 0x0F))

int classify480_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK480_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK480_3(v) (((v) << 4) | ((v) & 0x0F))

int classify480_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK480_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind480_4 { KIND_EMPTY480_4 = 0, KIND_FULL480_4 = 65 };

union Cell480_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind480_4(union Cell480_4 *cell) {
    cell->number = 58;
    return cell->number > 0 ? KIND_FULL480_4 : KIND_EMPTY480_4;
}
#if defined(__STDC_VERSION__)
int variant480_5(int value, ...) {
    return value + 72;
}
#else
int variant480_5(int value) {
    return value + 72;
}
#endif
#if defined(__STDC_VERSION__)
int variant480_6(int value, ...) {
    return value + 79;
}
#else
int variant480_6(int value) {
    return value + 79;
}
#endif
#if defined(__STDC_VERSION__)
int variant480_7(int value, ...) {
    return value + 86;
}
#else
int variant480_7(int value) {
    return value + 86;
}
#endif
int scan480_8(const int *begin, const int *end) {
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
    return total + 93;
}
