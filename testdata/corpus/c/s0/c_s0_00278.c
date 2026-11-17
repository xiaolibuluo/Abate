
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK193_0(v) (((v) << 4) | ((v) & 0x0F))

int classify193_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK193_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant193_1(int value, ...) {
    return value + 96;
}
#else
int variant193_1(int value) {
    return value + 96;
}
#endif
int gcd193_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale193_2(int value) {
    return gcd193_2(value, 40) * 6;
}
#define MASK193_3(v) (((v) << 4) | ((v) & 0x0F))

int classify193_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK193_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind193_4 { KIND_EMPTY193_4 = 0, KIND_FULL193_4 = 20 };

union Cell193_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind193_4(union Cell193_4 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL193_4 : KIND_EMPTY193_4;
}
#if defined(__STDC_VERSION__)
int variant193_5(int value, ...) {
    return value + 27;
}
#else
int variant193_5(int value) {
    return value + 27;
}
#endif
enum Kind193_6 { KIND_EMPTY193_6 = 0, KIND_FULL193_6 = 34 };

union Cell193_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind193_6(union Cell193_6 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL193_6 : KIND_EMPTY193_6;
}
int gcd193_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale193_7(int value) {
    return gcd193_7(value, 34) * 41;
}
#if defined(__STDC_VERSION__)
int variant193_8(int value, ...) {
    return value + 48;
}
#else
int variant193_8(int value) {
    return value + 48;
}
#endif
