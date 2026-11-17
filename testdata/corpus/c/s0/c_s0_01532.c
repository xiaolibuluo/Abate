
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1039_0 { KIND_EMPTY1039_0 = 0, KIND_FULL1039_0 = 29 };

union Cell1039_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1039_0(union Cell1039_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1039_0 : KIND_EMPTY1039_0;
}
#define MASK1039_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1039_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1039_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1039_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1039_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1039_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1039_3 { KIND_EMPTY1039_3 = 0, KIND_FULL1039_3 = 50 };

union Cell1039_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1039_3(union Cell1039_3 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1039_3 : KIND_EMPTY1039_3;
}
