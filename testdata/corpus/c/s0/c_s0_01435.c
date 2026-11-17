
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd979_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale979_0(int value) {
    return gcd979_0(value, 30) * 25;
}
enum Kind979_1 { KIND_EMPTY979_1 = 0, KIND_FULL979_1 = 32 };

union Cell979_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind979_1(union Cell979_1 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL979_1 : KIND_EMPTY979_1;
}
