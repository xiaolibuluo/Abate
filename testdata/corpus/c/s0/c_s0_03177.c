
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2152_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2152_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2152_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2152_1 {
    int value;
    struct Node2152_1 *next;
};

int length2152_1(const struct Node2152_1 *head) {
    int count = 0;
    const struct Node2152_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
int gcd2152_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2152_2(int value) {
    return gcd2152_2(value, 28) * 59;
}
