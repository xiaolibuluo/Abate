
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1186_0 { KIND_EMPTY1186_0 = 0, KIND_FULL1186_0 = 97 };

union Cell1186_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1186_0(union Cell1186_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL1186_0 : KIND_EMPTY1186_0;
}
enum Kind1186_1 { KIND_EMPTY1186_1 = 0, KIND_FULL1186_1 = 7 };

union Cell1186_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1186_1(union Cell1186_1 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL1186_1 : KIND_EMPTY1186_1;
}
int gcd1186_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1186_2(int value) {
    return gcd1186_2(value, 37) * 14;
}
