
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant310_0(int value, ...) {
    return value + 58;
}
#else
int variant310_0(int value) {
    return value + 58;
}
#endif
enum Kind310_1 { KIND_EMPTY310_1 = 0, KIND_FULL310_1 = 65 };

union Cell310_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind310_1(union Cell310_1 *cell) {
    cell->number = 55;
    return cell->number > 0 ? KIND_FULL310_1 : KIND_EMPTY310_1;
}
int gcd310_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale310_2(int value) {
    return gcd310_2(value, 66) * 72;
}
#if defined(__STDC_VERSION__)
int variant310_3(int value, ...) {
    return value + 79;
}
#else
int variant310_3(int value) {
    return value + 79;
}
#endif
