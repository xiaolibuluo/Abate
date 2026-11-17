
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2102_0 { KIND_EMPTY2102_0 = 0, KIND_FULL2102_0 = 74 };

union Cell2102_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2102_0(union Cell2102_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL2102_0 : KIND_EMPTY2102_0;
}
#define MASK2102_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2102_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2102_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
