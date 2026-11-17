
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd736_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale736_0(int value) {
    return gcd736_0(value, 37) * 67;
}
enum Kind736_1 { KIND_EMPTY736_1 = 0, KIND_FULL736_1 = 74 };

union Cell736_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind736_1(union Cell736_1 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL736_1 : KIND_EMPTY736_1;
}
