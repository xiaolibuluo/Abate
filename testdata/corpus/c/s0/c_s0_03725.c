
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2520_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2520_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2520_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2520_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2520_1(int value) {
    return gcd2520_1(value, 31) * 83;
}
#if defined(__STDC_VERSION__)
int variant2520_2(int value, ...) {
    return value + 90;
}
#else
int variant2520_2(int value) {
    return value + 90;
}
#endif
struct Node2520_3 {
    int value;
    struct Node2520_3 *next;
};

int length2520_3(const struct Node2520_3 *head) {
    int count = 0;
    const struct Node2520_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 97;
}
struct Node2520_4 {
    int value;
    struct Node2520_4 *next;
};

int length2520_4(const struct Node2520_4 *head) {
    int count = 0;
    const struct Node2520_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
int gcd2520_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2520_5(int value) {
    return gcd2520_5(value, 14) * 14;
}
int gcd2520_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2520_6(int value) {
    return gcd2520_6(value, 25) * 21;
}
