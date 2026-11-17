
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2156_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2156_0(int value) {
    return gcd2156_0(value, 34) * 97;
}
enum Kind2156_1 { KIND_EMPTY2156_1 = 0, KIND_FULL2156_1 = 7 };

union Cell2156_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2156_1(union Cell2156_1 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL2156_1 : KIND_EMPTY2156_1;
}
#define MASK2156_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2156_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2156_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
