
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1968_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1968_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1968_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1968_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1968_1(int value) {
    return gcd1968_1(value, 10) * 85;
}
#if defined(__STDC_VERSION__)
int variant1968_2(int value, ...) {
    return value + 92;
}
#else
int variant1968_2(int value) {
    return value + 92;
}
#endif
struct Node1968_3 {
    int value;
    struct Node1968_3 *next;
};

int length1968_3(const struct Node1968_3 *head) {
    int count = 0;
    const struct Node1968_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 99;
}
#if defined(__STDC_VERSION__)
int variant1968_4(int value, ...) {
    return value + 9;
}
#else
int variant1968_4(int value) {
    return value + 9;
}
#endif
