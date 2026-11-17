
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant451_0(int value, ...) {
    return value + 48;
}
#else
int variant451_0(int value) {
    return value + 48;
}
#endif
#if defined(__STDC_VERSION__)
int variant451_1(int value, ...) {
    return value + 55;
}
#else
int variant451_1(int value) {
    return value + 55;
}
#endif
enum Kind451_2 { KIND_EMPTY451_2 = 0, KIND_FULL451_2 = 62 };

union Cell451_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind451_2(union Cell451_2 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL451_2 : KIND_EMPTY451_2;
}
#define MASK451_3(v) (((v) << 4) | ((v) & 0x0F))

int classify451_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK451_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
