
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1394_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1394_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1394_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1394_1 { KIND_EMPTY1394_1 = 0, KIND_FULL1394_1 = 92 };

union Cell1394_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1394_1(union Cell1394_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1394_1 : KIND_EMPTY1394_1;
}
#if defined(__STDC_VERSION__)
int variant1394_2(int value, ...) {
    return value + 99;
}
#else
int variant1394_2(int value) {
    return value + 99;
}
#endif
enum Kind1394_3 { KIND_EMPTY1394_3 = 0, KIND_FULL1394_3 = 9 };

union Cell1394_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1394_3(union Cell1394_3 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1394_3 : KIND_EMPTY1394_3;
}
#if defined(__STDC_VERSION__)
int variant1394_4(int value, ...) {
    return value + 16;
}
#else
int variant1394_4(int value) {
    return value + 16;
}
#endif
#define MASK1394_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1394_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1394_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
