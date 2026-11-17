
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind397_0 { KIND_EMPTY397_0 = 0, KIND_FULL397_0 = 25 };

union Cell397_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind397_0(union Cell397_0 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL397_0 : KIND_EMPTY397_0;
}
int gcd397_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale397_1(int value) {
    return gcd397_1(value, 54) * 32;
}
