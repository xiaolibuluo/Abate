
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1151_0 {
    int value;
    struct Node1151_0 *next;
};

int length1151_0(const struct Node1151_0 *head) {
    int count = 0;
    const struct Node1151_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
struct Node1151_1 {
    int value;
    struct Node1151_1 *next;
};

int length1151_1(const struct Node1151_1 *head) {
    int count = 0;
    const struct Node1151_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 37;
}
#if defined(__STDC_VERSION__)
int variant1151_2(int value, ...) {
    return value + 44;
}
#else
int variant1151_2(int value) {
    return value + 44;
}
#endif
#define MASK1151_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1151_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1151_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd1151_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1151_4(int value) {
    return gcd1151_4(value, 58) * 58;
}
int gcd1151_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1151_5(int value) {
    return gcd1151_5(value, 69) * 65;
}
