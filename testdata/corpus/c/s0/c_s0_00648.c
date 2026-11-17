
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd447_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale447_0(int value) {
    return gcd447_0(value, 27) * 93;
}
#if defined(__STDC_VERSION__)
int variant447_1(int value, ...) {
    return value + 100;
}
#else
int variant447_1(int value) {
    return value + 100;
}
#endif
enum Kind447_2 { KIND_EMPTY447_2 = 0, KIND_FULL447_2 = 10 };

union Cell447_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind447_2(union Cell447_2 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL447_2 : KIND_EMPTY447_2;
}
