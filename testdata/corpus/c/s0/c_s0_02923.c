
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1978_0(int value, ...) {
    return value + 14;
}
#else
int variant1978_0(int value) {
    return value + 14;
}
#endif
#define MASK1978_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1978_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1978_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1978_2 {
    int value;
    struct Node1978_2 *next;
};

int length1978_2(const struct Node1978_2 *head) {
    int count = 0;
    const struct Node1978_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 28;
}
int gcd1978_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1978_3(int value) {
    return gcd1978_3(value, 41) * 35;
}
#if defined(__STDC_VERSION__)
int variant1978_4(int value, ...) {
    return value + 42;
}
#else
int variant1978_4(int value) {
    return value + 42;
}
#endif
struct Node1978_5 {
    int value;
    struct Node1978_5 *next;
};

int length1978_5(const struct Node1978_5 *head) {
    int count = 0;
    const struct Node1978_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 49;
}
