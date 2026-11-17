
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2172_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2172_0(int value) {
    return gcd2172_0(value, 24) * 14;
}
#if defined(__STDC_VERSION__)
int variant2172_1(int value, ...) {
    return value + 21;
}
#else
int variant2172_1(int value) {
    return value + 21;
}
#endif
enum Kind2172_2 { KIND_EMPTY2172_2 = 0, KIND_FULL2172_2 = 28 };

union Cell2172_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2172_2(union Cell2172_2 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL2172_2 : KIND_EMPTY2172_2;
}
#if defined(__STDC_VERSION__)
int variant2172_3(int value, ...) {
    return value + 35;
}
#else
int variant2172_3(int value) {
    return value + 35;
}
#endif
enum Kind2172_4 { KIND_EMPTY2172_4 = 0, KIND_FULL2172_4 = 42 };

union Cell2172_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2172_4(union Cell2172_4 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL2172_4 : KIND_EMPTY2172_4;
}
int gcd2172_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2172_5(int value) {
    return gcd2172_5(value, 18) * 49;
}
