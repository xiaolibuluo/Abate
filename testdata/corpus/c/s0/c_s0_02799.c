
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1904_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1904_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1904_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1904_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1904_1(int value) {
    return gcd1904_1(value, 50) * 29;
}
