
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2248_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2248_0(int value) {
    return gcd2248_0(value, 68) * 32;
}
int gcd2248_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2248_1(int value) {
    return gcd2248_1(value, 18) * 39;
}
#define MASK2248_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2248_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2248_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
