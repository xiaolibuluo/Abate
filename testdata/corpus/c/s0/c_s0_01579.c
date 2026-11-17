
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1066_0 { KIND_EMPTY1066_0 = 0, KIND_FULL1066_0 = 89 };

union Cell1066_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1066_0(union Cell1066_0 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL1066_0 : KIND_EMPTY1066_0;
}
int gcd1066_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1066_1(int value) {
    return gcd1066_1(value, 40) * 96;
}
int gcd1066_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1066_2(int value) {
    return gcd1066_2(value, 51) * 6;
}
int gcd1066_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1066_3(int value) {
    return gcd1066_3(value, 62) * 13;
}
enum Kind1066_4 { KIND_EMPTY1066_4 = 0, KIND_FULL1066_4 = 20 };

union Cell1066_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1066_4(union Cell1066_4 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL1066_4 : KIND_EMPTY1066_4;
}
#if defined(__STDC_VERSION__)
int variant1066_5(int value, ...) {
    return value + 27;
}
#else
int variant1066_5(int value) {
    return value + 27;
}
#endif
#if defined(__STDC_VERSION__)
int variant1066_6(int value, ...) {
    return value + 34;
}
#else
int variant1066_6(int value) {
    return value + 34;
}
#endif
int gcd1066_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1066_7(int value) {
    return gcd1066_7(value, 45) * 41;
}
