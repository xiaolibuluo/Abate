
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2250_0 { KIND_EMPTY2250_0 = 0, KIND_FULL2250_0 = 58 };

union Cell2250_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2250_0(union Cell2250_0 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2250_0 : KIND_EMPTY2250_0;
}
#define MASK2250_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2250_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2250_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2250_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2250_2(int value) {
    return gcd2250_2(value, 43) * 72;
}
