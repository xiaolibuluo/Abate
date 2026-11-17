
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2651_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2651_0(int value) {
    return gcd2651_0(value, 22) * 33;
}
struct Node2651_1 {
    int value;
    struct Node2651_1 *next;
};

int length2651_1(const struct Node2651_1 *head) {
    int count = 0;
    const struct Node2651_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 40;
}
#define MASK2651_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2651_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2651_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2651_3 {
    int value;
    struct Node2651_3 *next;
};

int length2651_3(const struct Node2651_3 *head) {
    int count = 0;
    const struct Node2651_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 54;
}
