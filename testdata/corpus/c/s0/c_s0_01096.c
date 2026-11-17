
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant749_0(int value, ...) {
    return value + 42;
}
#else
int variant749_0(int value) {
    return value + 42;
}
#endif
#define MASK749_1(v) (((v) << 4) | ((v) & 0x0F))

int classify749_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK749_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK749_2(v) (((v) << 4) | ((v) & 0x0F))

int classify749_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK749_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind749_3 { KIND_EMPTY749_3 = 0, KIND_FULL749_3 = 63 };

union Cell749_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind749_3(union Cell749_3 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL749_3 : KIND_EMPTY749_3;
}
