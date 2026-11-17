
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant252_0(int value, ...) {
    return value + 80;
}
#else
int variant252_0(int value) {
    return value + 80;
}
#endif
#define MASK252_1(v) (((v) << 4) | ((v) & 0x0F))

int classify252_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK252_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind252_2 { KIND_EMPTY252_2 = 0, KIND_FULL252_2 = 94 };

union Cell252_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind252_2(union Cell252_2 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL252_2 : KIND_EMPTY252_2;
}
