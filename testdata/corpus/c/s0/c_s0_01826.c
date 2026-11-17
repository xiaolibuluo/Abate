
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1226_0 { KIND_EMPTY1226_0 = 0, KIND_FULL1226_0 = 35 };

union Cell1226_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1226_0(union Cell1226_0 *cell) {
    cell->number = 51;
    return cell->number > 0 ? KIND_FULL1226_0 : KIND_EMPTY1226_0;
}
#if defined(__STDC_VERSION__)
int variant1226_1(int value, ...) {
    return value + 42;
}
#else
int variant1226_1(int value) {
    return value + 42;
}
#endif
int gcd1226_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1226_2(int value) {
    return gcd1226_2(value, 12) * 49;
}
