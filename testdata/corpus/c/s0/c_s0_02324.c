
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1568_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1568_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1568_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1568_1 { KIND_EMPTY1568_1 = 0, KIND_FULL1568_1 = 26 };

union Cell1568_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1568_1(union Cell1568_1 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1568_1 : KIND_EMPTY1568_1;
}
enum Kind1568_2 { KIND_EMPTY1568_2 = 0, KIND_FULL1568_2 = 33 };

union Cell1568_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1568_2(union Cell1568_2 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL1568_2 : KIND_EMPTY1568_2;
}
