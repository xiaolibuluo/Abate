
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1780_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1780_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1780_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1780_1(int value, ...) {
    return value + 66;
}
#else
int variant1780_1(int value) {
    return value + 66;
}
#endif
struct Node1780_2 {
    int value;
    struct Node1780_2 *next;
};

int length1780_2(const struct Node1780_2 *head) {
    int count = 0;
    const struct Node1780_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
struct Node1780_3 {
    int value;
    struct Node1780_3 *next;
};

int length1780_3(const struct Node1780_3 *head) {
    int count = 0;
    const struct Node1780_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
int gcd1780_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1780_4(int value) {
    return gcd1780_4(value, 69) * 87;
}
