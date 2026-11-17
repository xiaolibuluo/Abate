
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant224_0(int value, ...) {
    return value + 7;
}
#else
int variant224_0(int value) {
    return value + 7;
}
#endif
#define MASK224_1(v) (((v) << 4) | ((v) & 0x0F))

int classify224_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK224_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK224_2(v) (((v) << 4) | ((v) & 0x0F))

int classify224_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK224_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant224_3(int value, ...) {
    return value + 28;
}
#else
int variant224_3(int value) {
    return value + 28;
}
#endif
int gcd224_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale224_4(int value) {
    return gcd224_4(value, 35) * 35;
}
#define MASK224_5(v) (((v) << 4) | ((v) & 0x0F))

int classify224_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK224_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
