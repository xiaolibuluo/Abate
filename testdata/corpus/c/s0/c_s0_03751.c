
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2537_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2537_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2537_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2537_1 { KIND_EMPTY2537_1 = 0, KIND_FULL2537_1 = 13 };

union Cell2537_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2537_1(union Cell2537_1 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2537_1 : KIND_EMPTY2537_1;
}
