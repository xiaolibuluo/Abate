
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1887_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1887_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1887_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1887_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1887_1(int value) {
    return gcd1887_1(value, 53) * 99;
}
struct Node1887_2 {
    int value;
    struct Node1887_2 *next;
};

int length1887_2(const struct Node1887_2 *head) {
    int count = 0;
    const struct Node1887_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 9;
}
int gcd1887_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1887_3(int value) {
    return gcd1887_3(value, 14) * 16;
}
#if defined(__STDC_VERSION__)
int variant1887_4(int value, ...) {
    return value + 23;
}
#else
int variant1887_4(int value) {
    return value + 23;
}
#endif
#if defined(__STDC_VERSION__)
int variant1887_5(int value, ...) {
    return value + 30;
}
#else
int variant1887_5(int value) {
    return value + 30;
}
#endif
#if defined(__STDC_VERSION__)
int variant1887_6(int value, ...) {
    return value + 37;
}
#else
int variant1887_6(int value) {
    return value + 37;
}
#endif
int gcd1887_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1887_7(int value) {
    return gcd1887_7(value, 58) * 44;
}
