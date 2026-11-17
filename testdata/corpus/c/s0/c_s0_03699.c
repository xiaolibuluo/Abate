
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2500_0 { KIND_EMPTY2500_0 = 0, KIND_FULL2500_0 = 10 };

union Cell2500_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2500_0(union Cell2500_0 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL2500_0 : KIND_EMPTY2500_0;
}
#if defined(__STDC_VERSION__)
int variant2500_1(int value, ...) {
    return value + 17;
}
#else
int variant2500_1(int value) {
    return value + 17;
}
#endif
#if defined(__STDC_VERSION__)
int variant2500_2(int value, ...) {
    return value + 24;
}
#else
int variant2500_2(int value) {
    return value + 24;
}
#endif
int gcd2500_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2500_3(int value) {
    return gcd2500_3(value, 35) * 31;
}
