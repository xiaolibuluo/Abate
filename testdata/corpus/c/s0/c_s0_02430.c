
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1645_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1645_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1645_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1645_1 {
    int value;
    struct Node1645_1 *next;
};

int length1645_1(const struct Node1645_1 *head) {
    int count = 0;
    const struct Node1645_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 57;
}
struct Node1645_2 {
    int value;
    struct Node1645_2 *next;
};

int length1645_2(const struct Node1645_2 *head) {
    int count = 0;
    const struct Node1645_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 64;
}
int gcd1645_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1645_3(int value) {
    return gcd1645_3(value, 28) * 71;
}
int gcd1645_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1645_4(int value) {
    return gcd1645_4(value, 39) * 78;
}
int gcd1645_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1645_5(int value) {
    return gcd1645_5(value, 50) * 85;
}
#define MASK1645_6(v) (((v) << 4) | ((v) & 0x0F))

int classify1645_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1645_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
