
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1490_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1490_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1490_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1490_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1490_1(int value) {
    return gcd1490_1(value, 19) * 79;
}
enum Kind1490_2 { KIND_EMPTY1490_2 = 0, KIND_FULL1490_2 = 86 };

union Cell1490_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1490_2(union Cell1490_2 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1490_2 : KIND_EMPTY1490_2;
}
int gcd1490_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1490_3(int value) {
    return gcd1490_3(value, 41) * 93;
}
int gcd1490_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1490_4(int value) {
    return gcd1490_4(value, 52) * 100;
}
#define MASK1490_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1490_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1490_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
