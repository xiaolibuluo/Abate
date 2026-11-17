
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK782_0(v) (((v) << 4) | ((v) & 0x0F))

int classify782_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK782_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind782_1 { KIND_EMPTY782_1 = 0, KIND_FULL782_1 = 90 };

union Cell782_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind782_1(union Cell782_1 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL782_1 : KIND_EMPTY782_1;
}
