
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1379_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1379_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1379_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1379_1 { KIND_EMPTY1379_1 = 0, KIND_FULL1379_1 = 91 };

union Cell1379_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1379_1(union Cell1379_1 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1379_1 : KIND_EMPTY1379_1;
}
#if defined(__STDC_VERSION__)
int variant1379_2(int value, ...) {
    return value + 98;
}
#else
int variant1379_2(int value) {
    return value + 98;
}
#endif
#if defined(__STDC_VERSION__)
int variant1379_3(int value, ...) {
    return value + 8;
}
#else
int variant1379_3(int value) {
    return value + 8;
}
#endif
