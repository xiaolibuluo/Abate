
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd789_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale789_0(int value) {
    return gcd789_0(value, 42) * 77;
}
enum Kind789_1 { KIND_EMPTY789_1 = 0, KIND_FULL789_1 = 84 };

union Cell789_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind789_1(union Cell789_1 *cell) {
    cell->number = 53;
    return cell->number > 0 ? KIND_FULL789_1 : KIND_EMPTY789_1;
}
