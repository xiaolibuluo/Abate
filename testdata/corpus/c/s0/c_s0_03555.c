
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2399_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2399_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2399_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2399_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2399_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2399_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2399_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2399_2(int value) {
    return gcd2399_2(value, 49) * 69;
}
