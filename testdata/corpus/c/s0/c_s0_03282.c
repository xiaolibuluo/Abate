
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2219_0 { KIND_EMPTY2219_0 = 0, KIND_FULL2219_0 = 43 };

union Cell2219_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2219_0(union Cell2219_0 *cell) {
    cell->number = 48;
    return cell->number > 0 ? KIND_FULL2219_0 : KIND_EMPTY2219_0;
}
#if defined(__STDC_VERSION__)
int variant2219_1(int value, ...) {
    return value + 50;
}
#else
int variant2219_1(int value) {
    return value + 50;
}
#endif
int gcd2219_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2219_2(int value) {
    return gcd2219_2(value, 9) * 57;
}
#if defined(__STDC_VERSION__)
int variant2219_3(int value, ...) {
    return value + 64;
}
#else
int variant2219_3(int value) {
    return value + 64;
}
#endif
int gcd2219_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2219_4(int value) {
    return gcd2219_4(value, 31) * 71;
}
#if defined(__STDC_VERSION__)
int variant2219_5(int value, ...) {
    return value + 78;
}
#else
int variant2219_5(int value) {
    return value + 78;
}
#endif
