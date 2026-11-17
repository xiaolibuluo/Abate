
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind839_0 { KIND_EMPTY839_0 = 0, KIND_FULL839_0 = 48 };

union Cell839_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind839_0(union Cell839_0 *cell) {
    cell->number = 26;
    return cell->number > 0 ? KIND_FULL839_0 : KIND_EMPTY839_0;
}
#if defined(__STDC_VERSION__)
int variant839_1(int value, ...) {
    return value + 55;
}
#else
int variant839_1(int value) {
    return value + 55;
}
#endif
int gcd839_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale839_2(int value) {
    return gcd839_2(value, 48) * 62;
}
