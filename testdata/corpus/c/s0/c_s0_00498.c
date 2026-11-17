
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd344_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale344_0(int value) {
    return gcd344_0(value, 38) * 15;
}
enum Kind344_1 { KIND_EMPTY344_1 = 0, KIND_FULL344_1 = 22 };

union Cell344_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind344_1(union Cell344_1 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL344_1 : KIND_EMPTY344_1;
}
