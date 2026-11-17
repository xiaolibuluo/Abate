
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1259_0 { KIND_EMPTY1259_0 = 0, KIND_FULL1259_0 = 76 };

union Cell1259_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1259_0(union Cell1259_0 *cell) {
    cell->number = 38;
    return cell->number > 0 ? KIND_FULL1259_0 : KIND_EMPTY1259_0;
}
int gcd1259_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1259_1(int value) {
    return gcd1259_1(value, 49) * 83;
}
enum Kind1259_2 { KIND_EMPTY1259_2 = 0, KIND_FULL1259_2 = 90 };

union Cell1259_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1259_2(union Cell1259_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL1259_2 : KIND_EMPTY1259_2;
}
int gcd1259_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1259_3(int value) {
    return gcd1259_3(value, 10) * 97;
}
int gcd1259_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1259_4(int value) {
    return gcd1259_4(value, 21) * 7;
}
enum Kind1259_5 { KIND_EMPTY1259_5 = 0, KIND_FULL1259_5 = 14 };

union Cell1259_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1259_5(union Cell1259_5 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1259_5 : KIND_EMPTY1259_5;
}
enum Kind1259_6 { KIND_EMPTY1259_6 = 0, KIND_FULL1259_6 = 21 };

union Cell1259_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1259_6(union Cell1259_6 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL1259_6 : KIND_EMPTY1259_6;
}
