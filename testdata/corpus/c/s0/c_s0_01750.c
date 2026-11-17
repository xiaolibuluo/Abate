
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1175_0 { KIND_EMPTY1175_0 = 0, KIND_FULL1175_0 = 51 };

union Cell1175_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1175_0(union Cell1175_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1175_0 : KIND_EMPTY1175_0;
}
int gcd1175_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1175_1(int value) {
    return gcd1175_1(value, 10) * 58;
}
