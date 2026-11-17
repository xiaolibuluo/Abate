
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1919_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1919_0(int value) {
    return gcd1919_0(value, 22) * 23;
}
#define MASK1919_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1919_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1919_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1919_2 { KIND_EMPTY1919_2 = 0, KIND_FULL1919_2 = 37 };

union Cell1919_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1919_2(union Cell1919_2 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1919_2 : KIND_EMPTY1919_2;
}
enum Kind1919_3 { KIND_EMPTY1919_3 = 0, KIND_FULL1919_3 = 44 };

union Cell1919_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1919_3(union Cell1919_3 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL1919_3 : KIND_EMPTY1919_3;
}
