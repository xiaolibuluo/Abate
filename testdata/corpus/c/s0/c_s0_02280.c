
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1540_0(int value, ...) {
    return value + 43;
}
#else
int variant1540_0(int value) {
    return value + 43;
}
#endif
enum Kind1540_1 { KIND_EMPTY1540_1 = 0, KIND_FULL1540_1 = 50 };

union Cell1540_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1540_1(union Cell1540_1 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL1540_1 : KIND_EMPTY1540_1;
}
int gcd1540_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1540_2(int value) {
    return gcd1540_2(value, 14) * 57;
}
int gcd1540_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1540_3(int value) {
    return gcd1540_3(value, 25) * 64;
}
enum Kind1540_4 { KIND_EMPTY1540_4 = 0, KIND_FULL1540_4 = 71 };

union Cell1540_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1540_4(union Cell1540_4 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL1540_4 : KIND_EMPTY1540_4;
}
enum Kind1540_5 { KIND_EMPTY1540_5 = 0, KIND_FULL1540_5 = 78 };

union Cell1540_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1540_5(union Cell1540_5 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL1540_5 : KIND_EMPTY1540_5;
}
