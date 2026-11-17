
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2306_0(int value, ...) {
    return value + 10;
}
#else
int variant2306_0(int value) {
    return value + 10;
}
#endif
#define MASK2306_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2306_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2306_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2306_2 {
    int value;
    struct Node2306_2 *next;
};

int length2306_2(const struct Node2306_2 *head) {
    int count = 0;
    const struct Node2306_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
int gcd2306_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2306_3(int value) {
    return gcd2306_3(value, 19) * 31;
}
