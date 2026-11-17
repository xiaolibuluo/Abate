
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1085_0 { KIND_EMPTY1085_0 = 0, KIND_FULL1085_0 = 45 };

union Cell1085_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1085_0(union Cell1085_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1085_0 : KIND_EMPTY1085_0;
}
int gcd1085_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1085_1(int value) {
    return gcd1085_1(value, 51) * 52;
}
#define MASK1085_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1085_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1085_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
