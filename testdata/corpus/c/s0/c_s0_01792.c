
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1202_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1202_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1202_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1202_1 { KIND_EMPTY1202_1 = 0, KIND_FULL1202_1 = 21 };

union Cell1202_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1202_1(union Cell1202_1 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1202_1 : KIND_EMPTY1202_1;
}
#define MASK1202_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1202_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1202_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1202_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1202_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1202_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1202_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1202_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1202_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1202_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1202_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1202_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
