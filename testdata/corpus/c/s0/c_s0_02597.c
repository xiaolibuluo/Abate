
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1760_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1760_0(int value) {
    return gcd1760_0(value, 68) * 90;
}
#if defined(__STDC_VERSION__)
int variant1760_1(int value, ...) {
    return value + 97;
}
#else
int variant1760_1(int value) {
    return value + 97;
}
#endif
enum Kind1760_2 { KIND_EMPTY1760_2 = 0, KIND_FULL1760_2 = 7 };

union Cell1760_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1760_2(union Cell1760_2 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1760_2 : KIND_EMPTY1760_2;
}
