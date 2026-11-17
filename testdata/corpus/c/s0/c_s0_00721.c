
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK497_0(v) (((v) << 4) | ((v) & 0x0F))

int classify497_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK497_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind497_1 { KIND_EMPTY497_1 = 0, KIND_FULL497_1 = 71 };

union Cell497_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind497_1(union Cell497_1 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL497_1 : KIND_EMPTY497_1;
}
enum Kind497_2 { KIND_EMPTY497_2 = 0, KIND_FULL497_2 = 78 };

union Cell497_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind497_2(union Cell497_2 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL497_2 : KIND_EMPTY497_2;
}
