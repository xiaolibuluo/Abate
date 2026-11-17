
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind834_0 { KIND_EMPTY834_0 = 0, KIND_FULL834_0 = 80 };

union Cell834_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind834_0(union Cell834_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL834_0 : KIND_EMPTY834_0;
}
int gcd834_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale834_1(int value) {
    return gcd834_1(value, 63) * 87;
}
