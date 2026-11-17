
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2397_0(int value, ...) {
    return value + 29;
}
#else
int variant2397_0(int value) {
    return value + 29;
}
#endif
int gcd2397_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2397_1(int value) {
    return gcd2397_1(value, 24) * 36;
}
enum Kind2397_2 { KIND_EMPTY2397_2 = 0, KIND_FULL2397_2 = 43 };

union Cell2397_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2397_2(union Cell2397_2 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL2397_2 : KIND_EMPTY2397_2;
}
