
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2257_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2257_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2257_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2257_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2257_1(int value) {
    return gcd2257_1(value, 20) * 59;
}
enum Kind2257_2 { KIND_EMPTY2257_2 = 0, KIND_FULL2257_2 = 66 };

union Cell2257_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2257_2(union Cell2257_2 *cell) {
    cell->number = 31;
    return cell->number > 0 ? KIND_FULL2257_2 : KIND_EMPTY2257_2;
}
#define MASK2257_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2257_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2257_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
