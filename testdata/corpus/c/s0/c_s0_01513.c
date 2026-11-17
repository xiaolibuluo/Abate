
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1028_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1028_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1028_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1028_1 { KIND_EMPTY1028_1 = 0, KIND_FULL1028_1 = 87 };

union Cell1028_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1028_1(union Cell1028_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1028_1 : KIND_EMPTY1028_1;
}
