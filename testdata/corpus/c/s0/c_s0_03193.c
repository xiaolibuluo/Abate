
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2162_0 { KIND_EMPTY2162_0 = 0, KIND_FULL2162_0 = 78 };

union Cell2162_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2162_0(union Cell2162_0 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL2162_0 : KIND_EMPTY2162_0;
}
int gcd2162_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2162_1(int value) {
    return gcd2162_1(value, 26) * 85;
}
enum Kind2162_2 { KIND_EMPTY2162_2 = 0, KIND_FULL2162_2 = 92 };

union Cell2162_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2162_2(union Cell2162_2 *cell) {
    cell->number = 37;
    return cell->number > 0 ? KIND_FULL2162_2 : KIND_EMPTY2162_2;
}
int gcd2162_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2162_3(int value) {
    return gcd2162_3(value, 48) * 99;
}
