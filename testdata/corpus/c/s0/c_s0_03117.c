
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2111_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2111_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2111_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK2111_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2111_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2111_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2111_2 {
    int value;
    struct Node2111_2 *next;
};

int length2111_2(const struct Node2111_2 *head) {
    int count = 0;
    const struct Node2111_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
int gcd2111_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2111_3(int value) {
    return gcd2111_3(value, 57) * 18;
}
