
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd298_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale298_0(int value) {
    return gcd298_0(value, 21) * 96;
}
enum Kind298_1 { KIND_EMPTY298_1 = 0, KIND_FULL298_1 = 6 };

union Cell298_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind298_1(union Cell298_1 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL298_1 : KIND_EMPTY298_1;
}
