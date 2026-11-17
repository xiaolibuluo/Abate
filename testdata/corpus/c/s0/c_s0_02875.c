
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1951_0 { KIND_EMPTY1951_0 = 0, KIND_FULL1951_0 = 51 };

union Cell1951_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1951_0(union Cell1951_0 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1951_0 : KIND_EMPTY1951_0;
}
#define MASK1951_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1951_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1951_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
