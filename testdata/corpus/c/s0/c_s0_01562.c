
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1054_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1054_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1054_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1054_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1054_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1054_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1054_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1054_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1054_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1054_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1054_3(int value) {
    return gcd1054_3(value, 39) * 51;
}
#define MASK1054_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1054_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1054_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1054_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1054_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1054_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
