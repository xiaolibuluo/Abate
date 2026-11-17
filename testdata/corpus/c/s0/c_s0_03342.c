
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2255_0 { KIND_EMPTY2255_0 = 0, KIND_FULL2255_0 = 26 };

union Cell2255_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2255_0(union Cell2255_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL2255_0 : KIND_EMPTY2255_0;
}
#define MASK2255_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2255_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2255_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2255_2 { KIND_EMPTY2255_2 = 0, KIND_FULL2255_2 = 40 };

union Cell2255_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2255_2(union Cell2255_2 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL2255_2 : KIND_EMPTY2255_2;
}
