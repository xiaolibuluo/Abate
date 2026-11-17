
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1579_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1579_0(int value) {
    return gcd1579_0(value, 21) * 65;
}
int gcd1579_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1579_1(int value) {
    return gcd1579_1(value, 32) * 72;
}
#if defined(__STDC_VERSION__)
int variant1579_2(int value, ...) {
    return value + 79;
}
#else
int variant1579_2(int value) {
    return value + 79;
}
#endif
#define MASK1579_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1579_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1579_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1579_4 {
    int value;
    struct Node1579_4 *next;
};

int length1579_4(const struct Node1579_4 *head) {
    int count = 0;
    const struct Node1579_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
