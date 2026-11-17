
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1636_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1636_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1636_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1636_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1636_1(int value) {
    return gcd1636_1(value, 65) * 37;
}
#if defined(__STDC_VERSION__)
int variant1636_2(int value, ...) {
    return value + 44;
}
#else
int variant1636_2(int value) {
    return value + 44;
}
#endif
#define MASK1636_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1636_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1636_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1636_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1636_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1636_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1636_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1636_5(int value) {
    return gcd1636_5(value, 48) * 65;
}
