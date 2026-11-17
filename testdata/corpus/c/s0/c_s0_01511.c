
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1027_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1027_0(int value) {
    return gcd1027_0(value, 61) * 67;
}
enum Kind1027_1 { KIND_EMPTY1027_1 = 0, KIND_FULL1027_1 = 74 };

union Cell1027_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1027_1(union Cell1027_1 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1027_1 : KIND_EMPTY1027_1;
}
