
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd103_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale103_0(int value) {
    return gcd103_0(value, 59) * 83;
}
enum Kind103_1 { KIND_EMPTY103_1 = 0, KIND_FULL103_1 = 90 };

union Cell103_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind103_1(union Cell103_1 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL103_1 : KIND_EMPTY103_1;
}
