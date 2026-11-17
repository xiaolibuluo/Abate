
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2358_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2358_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2358_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2358_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2358_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2358_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2358_2 { KIND_EMPTY2358_2 = 0, KIND_FULL2358_2 = 21 };

union Cell2358_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2358_2(union Cell2358_2 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2358_2 : KIND_EMPTY2358_2;
}
#define MASK2358_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2358_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2358_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2358_4 { KIND_EMPTY2358_4 = 0, KIND_FULL2358_4 = 35 };

union Cell2358_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2358_4(union Cell2358_4 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2358_4 : KIND_EMPTY2358_4;
}
