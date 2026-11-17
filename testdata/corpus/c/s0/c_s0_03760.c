
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2544_0(int value, ...) {
    return value + 97;
}
#else
int variant2544_0(int value) {
    return value + 97;
}
#endif
int gcd2544_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2544_1(int value) {
    return gcd2544_1(value, 16) * 7;
}
enum Kind2544_2 { KIND_EMPTY2544_2 = 0, KIND_FULL2544_2 = 14 };

union Cell2544_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2544_2(union Cell2544_2 *cell) {
    cell->number = 27;
    return cell->number > 0 ? KIND_FULL2544_2 : KIND_EMPTY2544_2;
}
