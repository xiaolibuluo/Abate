
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1380_0 { KIND_EMPTY1380_0 = 0, KIND_FULL1380_0 = 97 };

union Cell1380_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1380_0(union Cell1380_0 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL1380_0 : KIND_EMPTY1380_0;
}
#define MASK1380_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1380_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1380_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1380_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1380_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1380_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1380_3 { KIND_EMPTY1380_3 = 0, KIND_FULL1380_3 = 21 };

union Cell1380_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1380_3(union Cell1380_3 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1380_3 : KIND_EMPTY1380_3;
}
