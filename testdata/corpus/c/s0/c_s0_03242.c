
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2197_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2197_0(int value) {
    return gcd2197_0(value, 16) * 48;
}
#define MASK2197_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2197_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2197_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
