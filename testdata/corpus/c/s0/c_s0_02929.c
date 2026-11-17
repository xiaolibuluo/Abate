
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1980_0 { KIND_EMPTY1980_0 = 0, KIND_FULL1980_0 = 40 };

union Cell1980_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1980_0(union Cell1980_0 *cell) {
    cell->number = 22;
    return cell->number > 0 ? KIND_FULL1980_0 : KIND_EMPTY1980_0;
}
enum Kind1980_1 { KIND_EMPTY1980_1 = 0, KIND_FULL1980_1 = 47 };

union Cell1980_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1980_1(union Cell1980_1 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL1980_1 : KIND_EMPTY1980_1;
}
enum Kind1980_2 { KIND_EMPTY1980_2 = 0, KIND_FULL1980_2 = 54 };

union Cell1980_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1980_2(union Cell1980_2 *cell) {
    cell->number = 44;
    return cell->number > 0 ? KIND_FULL1980_2 : KIND_EMPTY1980_2;
}
#define MASK1980_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1980_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1980_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1980_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1980_4(int value) {
    return gcd1980_4(value, 66) * 68;
}
int gcd1980_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1980_5(int value) {
    return gcd1980_5(value, 16) * 75;
}
