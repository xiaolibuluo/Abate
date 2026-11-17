
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1274_0 { KIND_EMPTY1274_0 = 0, KIND_FULL1274_0 = 77 };

union Cell1274_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1274_0(union Cell1274_0 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL1274_0 : KIND_EMPTY1274_0;
}
int gcd1274_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1274_1(int value) {
    return gcd1274_1(value, 32) * 84;
}
