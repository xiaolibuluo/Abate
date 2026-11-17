
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant808_0(int value, ...) {
    return value + 33;
}
#else
int variant808_0(int value) {
    return value + 33;
}
#endif
int gcd808_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale808_1(int value) {
    return gcd808_1(value, 64) * 40;
}
#if defined(__STDC_VERSION__)
int variant808_2(int value, ...) {
    return value + 47;
}
#else
int variant808_2(int value) {
    return value + 47;
}
#endif
#define MASK808_3(v) (((v) << 4) | ((v) & 0x0F))

int classify808_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK808_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd808_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale808_4(int value) {
    return gcd808_4(value, 36) * 61;
}
enum Kind808_5 { KIND_EMPTY808_5 = 0, KIND_FULL808_5 = 68 };

union Cell808_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind808_5(union Cell808_5 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL808_5 : KIND_EMPTY808_5;
}
