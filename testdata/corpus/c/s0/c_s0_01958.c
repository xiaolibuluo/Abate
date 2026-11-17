
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1321_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1321_0(int value) {
    return gcd1321_0(value, 45) * 9;
}
enum Kind1321_1 { KIND_EMPTY1321_1 = 0, KIND_FULL1321_1 = 16 };

union Cell1321_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1321_1(union Cell1321_1 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL1321_1 : KIND_EMPTY1321_1;
}
