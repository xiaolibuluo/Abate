
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK700_0(v) (((v) << 4) | ((v) & 0x0F))

int classify700_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK700_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK700_1(v) (((v) << 4) | ((v) & 0x0F))

int classify700_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK700_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind700_2 { KIND_EMPTY700_2 = 0, KIND_FULL700_2 = 98 };

union Cell700_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind700_2(union Cell700_2 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL700_2 : KIND_EMPTY700_2;
}
enum Kind700_3 { KIND_EMPTY700_3 = 0, KIND_FULL700_3 = 8 };

union Cell700_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind700_3(union Cell700_3 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL700_3 : KIND_EMPTY700_3;
}
