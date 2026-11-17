
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1863_0(int value, ...) {
    return value + 71;
}
#else
int variant1863_0(int value) {
    return value + 71;
}
#endif
int gcd1863_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1863_1(int value) {
    return gcd1863_1(value, 68) * 78;
}
enum Kind1863_2 { KIND_EMPTY1863_2 = 0, KIND_FULL1863_2 = 85 };

union Cell1863_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1863_2(union Cell1863_2 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL1863_2 : KIND_EMPTY1863_2;
}
