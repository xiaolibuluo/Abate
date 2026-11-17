
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1713_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1713_0(int value) {
    return gcd1713_0(value, 44) * 61;
}
#if defined(__STDC_VERSION__)
int variant1713_1(int value, ...) {
    return value + 68;
}
#else
int variant1713_1(int value) {
    return value + 68;
}
#endif
enum Kind1713_2 { KIND_EMPTY1713_2 = 0, KIND_FULL1713_2 = 75 };

union Cell1713_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1713_2(union Cell1713_2 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1713_2 : KIND_EMPTY1713_2;
}
