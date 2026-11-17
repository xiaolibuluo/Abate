
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2440_0 { KIND_EMPTY2440_0 = 0, KIND_FULL2440_0 = 6 };

union Cell2440_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2440_0(union Cell2440_0 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL2440_0 : KIND_EMPTY2440_0;
}
int gcd2440_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2440_1(int value) {
    return gcd2440_1(value, 20) * 13;
}
