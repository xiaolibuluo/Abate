
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2155_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2155_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2155_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2155_1 {
    int value;
    struct Node2155_1 *next;
};

int length2155_1(const struct Node2155_1 *head) {
    int count = 0;
    const struct Node2155_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 91;
}
int gcd2155_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2155_2(int value) {
    return gcd2155_2(value, 49) * 98;
}
