
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1150_0(int value, ...) {
    return value + 17;
}
#else
int variant1150_0(int value) {
    return value + 17;
}
#endif
#define MASK1150_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1150_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1150_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1150_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1150_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1150_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1150_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1150_3(int value) {
    return gcd1150_3(value, 40) * 38;
}
int gcd1150_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1150_4(int value) {
    return gcd1150_4(value, 51) * 45;
}
#define MASK1150_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1150_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1150_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
