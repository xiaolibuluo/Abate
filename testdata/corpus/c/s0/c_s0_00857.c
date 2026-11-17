
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd590_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale590_0(int value) {
    return gcd590_0(value, 52) * 12;
}
enum Kind590_1 { KIND_EMPTY590_1 = 0, KIND_FULL590_1 = 19 };

union Cell590_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind590_1(union Cell590_1 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL590_1 : KIND_EMPTY590_1;
}
