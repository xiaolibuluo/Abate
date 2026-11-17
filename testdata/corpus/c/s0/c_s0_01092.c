
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant747_0(int value, ...) {
    return value + 16;
}
#else
int variant747_0(int value) {
    return value + 16;
}
#endif
#define MASK747_1(v) (((v) << 4) | ((v) & 0x0F))

int classify747_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK747_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK747_2(v) (((v) << 4) | ((v) & 0x0F))

int classify747_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK747_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd747_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale747_3(int value) {
    return gcd747_3(value, 25) * 37;
}
