
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK547_0(v) (((v) << 4) | ((v) & 0x0F))

int classify547_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK547_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind547_1 { KIND_EMPTY547_1 = 0, KIND_FULL547_1 = 42 };

union Cell547_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind547_1(union Cell547_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL547_1 : KIND_EMPTY547_1;
}
#define MASK547_2(v) (((v) << 4) | ((v) & 0x0F))

int classify547_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK547_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
