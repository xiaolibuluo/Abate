
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd676_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale676_0(int value) {
    return gcd676_0(value, 44) * 63;
}
int gcd676_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale676_1(int value) {
    return gcd676_1(value, 55) * 70;
}
enum Kind676_2 { KIND_EMPTY676_2 = 0, KIND_FULL676_2 = 77 };

union Cell676_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind676_2(union Cell676_2 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL676_2 : KIND_EMPTY676_2;
}
enum Kind676_3 { KIND_EMPTY676_3 = 0, KIND_FULL676_3 = 84 };

union Cell676_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind676_3(union Cell676_3 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL676_3 : KIND_EMPTY676_3;
}
#if defined(__STDC_VERSION__)
int variant676_4(int value, ...) {
    return value + 91;
}
#else
int variant676_4(int value) {
    return value + 91;
}
#endif
int scan676_5(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 98;
}
enum Kind676_6 { KIND_EMPTY676_6 = 0, KIND_FULL676_6 = 8 };

union Cell676_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind676_6(union Cell676_6 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL676_6 : KIND_EMPTY676_6;
}
