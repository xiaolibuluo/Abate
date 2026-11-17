
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1466_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1466_0(int value) {
    return gcd1466_0(value, 23) * 51;
}
enum Kind1466_1 { KIND_EMPTY1466_1 = 0, KIND_FULL1466_1 = 58 };

union Cell1466_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1466_1(union Cell1466_1 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1466_1 : KIND_EMPTY1466_1;
}
