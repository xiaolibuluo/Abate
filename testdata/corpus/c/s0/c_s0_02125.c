
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1432_0(int value, ...) {
    return value + 94;
}
#else
int variant1432_0(int value) {
    return value + 94;
}
#endif
enum Kind1432_1 { KIND_EMPTY1432_1 = 0, KIND_FULL1432_1 = 101 };

union Cell1432_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1432_1(union Cell1432_1 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1432_1 : KIND_EMPTY1432_1;
}
int gcd1432_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1432_2(int value) {
    return gcd1432_2(value, 51) * 11;
}
enum Kind1432_3 { KIND_EMPTY1432_3 = 0, KIND_FULL1432_3 = 18 };

union Cell1432_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1432_3(union Cell1432_3 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL1432_3 : KIND_EMPTY1432_3;
}
