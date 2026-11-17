
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK609_0(v) (((v) << 4) | ((v) & 0x0F))

int classify609_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK609_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant609_1(int value, ...) {
    return value + 72;
}
#else
int variant609_1(int value) {
    return value + 72;
}
#endif
enum Kind609_2 { KIND_EMPTY609_2 = 0, KIND_FULL609_2 = 79 };

union Cell609_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind609_2(union Cell609_2 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL609_2 : KIND_EMPTY609_2;
}
int gcd609_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale609_3(int value) {
    return gcd609_3(value, 35) * 86;
}
#if defined(__STDC_VERSION__)
int variant609_4(int value, ...) {
    return value + 93;
}
#else
int variant609_4(int value) {
    return value + 93;
}
#endif
#define MASK609_5(v) (((v) << 4) | ((v) & 0x0F))

int classify609_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK609_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
