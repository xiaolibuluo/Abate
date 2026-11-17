
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1272_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1272_0(int value) {
    return gcd1272_0(value, 68) * 51;
}
enum Kind1272_1 { KIND_EMPTY1272_1 = 0, KIND_FULL1272_1 = 58 };

union Cell1272_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1272_1(union Cell1272_1 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1272_1 : KIND_EMPTY1272_1;
}
