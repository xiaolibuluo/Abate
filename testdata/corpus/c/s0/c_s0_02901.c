
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1967_0(int value, ...) {
    return value + 65;
}
#else
int variant1967_0(int value) {
    return value + 65;
}
#endif
struct Node1967_1 {
    int value;
    struct Node1967_1 *next;
};

int length1967_1(const struct Node1967_1 *head) {
    int count = 0;
    const struct Node1967_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 72;
}
#if defined(__STDC_VERSION__)
int variant1967_2(int value, ...) {
    return value + 79;
}
#else
int variant1967_2(int value) {
    return value + 79;
}
#endif
#define MASK1967_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1967_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1967_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1967_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1967_4(int value) {
    return gcd1967_4(value, 36) * 93;
}
