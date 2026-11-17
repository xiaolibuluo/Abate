
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind498_0 { KIND_EMPTY498_0 = 0, KIND_FULL498_0 = 77 };

union Cell498_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind498_0(union Cell498_0 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL498_0 : KIND_EMPTY498_0;
}
#if defined(__STDC_VERSION__)
int variant498_1(int value, ...) {
    return value + 84;
}
#else
int variant498_1(int value) {
    return value + 84;
}
#endif
int gcd498_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale498_2(int value) {
    return gcd498_2(value, 40) * 91;
}
