
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1033_0(int value, ...) {
    return value + 48;
}
#else
int variant1033_0(int value) {
    return value + 48;
}
#endif
int gcd1033_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1033_1(int value) {
    return gcd1033_1(value, 53) * 55;
}
enum Kind1033_2 { KIND_EMPTY1033_2 = 0, KIND_FULL1033_2 = 62 };

union Cell1033_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1033_2(union Cell1033_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1033_2 : KIND_EMPTY1033_2;
}
