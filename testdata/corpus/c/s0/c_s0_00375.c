
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant253_0(int value, ...) {
    return value + 93;
}
#else
int variant253_0(int value) {
    return value + 93;
}
#endif
int gcd253_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale253_1(int value) {
    return gcd253_1(value, 22) * 100;
}
enum Kind253_2 { KIND_EMPTY253_2 = 0, KIND_FULL253_2 = 10 };

union Cell253_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind253_2(union Cell253_2 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL253_2 : KIND_EMPTY253_2;
}
