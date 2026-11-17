
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd217_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale217_0(int value) {
    return gcd217_0(value, 64) * 13;
}
int gcd217_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale217_1(int value) {
    return gcd217_1(value, 14) * 20;
}
#define MASK217_2(v) (((v) << 4) | ((v) & 0x0F))

int classify217_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK217_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd217_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale217_3(int value) {
    return gcd217_3(value, 36) * 34;
}
enum Kind217_4 { KIND_EMPTY217_4 = 0, KIND_FULL217_4 = 41 };

union Cell217_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind217_4(union Cell217_4 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL217_4 : KIND_EMPTY217_4;
}
