
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1224_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1224_0(int value) {
    return gcd1224_0(value, 37) * 9;
}
enum Kind1224_1 { KIND_EMPTY1224_1 = 0, KIND_FULL1224_1 = 16 };

union Cell1224_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1224_1(union Cell1224_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL1224_1 : KIND_EMPTY1224_1;
}
