
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1222_0 { KIND_EMPTY1222_0 = 0, KIND_FULL1222_0 = 80 };

union Cell1222_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1222_0(union Cell1222_0 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1222_0 : KIND_EMPTY1222_0;
}
int gcd1222_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1222_1(int value) {
    return gcd1222_1(value, 34) * 87;
}
