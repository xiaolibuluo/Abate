
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind493_0 { KIND_EMPTY493_0 = 0, KIND_FULL493_0 = 12 };

union Cell493_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind493_0(union Cell493_0 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL493_0 : KIND_EMPTY493_0;
}
#define MASK493_1(v) (((v) << 4) | ((v) & 0x0F))

int classify493_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK493_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
