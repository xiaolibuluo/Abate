
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2603_0 {
    int value;
    struct Node2603_0 *next;
};

int length2603_0(const struct Node2603_0 *head) {
    int count = 0;
    const struct Node2603_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
#define MASK2603_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2603_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2603_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2603_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2603_2(int value) {
    return gcd2603_2(value, 13) * 5;
}
struct Node2603_3 {
    int value;
    struct Node2603_3 *next;
};

int length2603_3(const struct Node2603_3 *head) {
    int count = 0;
    const struct Node2603_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
