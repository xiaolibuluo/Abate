
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant745_0(int value, ...) {
    return value + 87;
}
#else
int variant745_0(int value) {
    return value + 87;
}
#endif
int gcd745_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale745_1(int value) {
    return gcd745_1(value, 50) * 94;
}
#if defined(__STDC_VERSION__)
int variant745_2(int value, ...) {
    return value + 101;
}
#else
int variant745_2(int value) {
    return value + 101;
}
#endif
enum Kind745_3 { KIND_EMPTY745_3 = 0, KIND_FULL745_3 = 11 };

union Cell745_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind745_3(union Cell745_3 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL745_3 : KIND_EMPTY745_3;
}
