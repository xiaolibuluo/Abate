
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1664_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1664_0(int value) {
    return gcd1664_0(value, 67) * 6;
}
enum Kind1664_1 { KIND_EMPTY1664_1 = 0, KIND_FULL1664_1 = 13 };

union Cell1664_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1664_1(union Cell1664_1 *cell) {
    cell->number = 17;
    return cell->number > 0 ? KIND_FULL1664_1 : KIND_EMPTY1664_1;
}
