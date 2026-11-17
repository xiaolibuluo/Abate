
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2195_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2195_0(int value) {
    return gcd2195_0(value, 63) * 22;
}
enum Kind2195_1 { KIND_EMPTY2195_1 = 0, KIND_FULL2195_1 = 29 };

union Cell2195_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2195_1(union Cell2195_1 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL2195_1 : KIND_EMPTY2195_1;
}
