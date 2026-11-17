
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1758_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1758_0(int value) {
    return gcd1758_0(value, 54) * 64;
}
enum Kind1758_1 { KIND_EMPTY1758_1 = 0, KIND_FULL1758_1 = 71 };

union Cell1758_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1758_1(union Cell1758_1 *cell) {
    cell->number = 65;
    return cell->number > 0 ? KIND_FULL1758_1 : KIND_EMPTY1758_1;
}
