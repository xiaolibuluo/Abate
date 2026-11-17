
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind548_0 { KIND_EMPTY548_0 = 0, KIND_FULL548_0 = 48 };

union Cell548_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind548_0(union Cell548_0 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL548_0 : KIND_EMPTY548_0;
}
enum Kind548_1 { KIND_EMPTY548_1 = 0, KIND_FULL548_1 = 55 };

union Cell548_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind548_1(union Cell548_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL548_1 : KIND_EMPTY548_1;
}
int gcd548_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale548_2(int value) {
    return gcd548_2(value, 24) * 62;
}
