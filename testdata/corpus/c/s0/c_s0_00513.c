
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK356_0(v) (((v) << 4) | ((v) & 0x0F))

int classify356_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK356_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd356_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale356_1(int value) {
    return gcd356_1(value, 11) * 81;
}
struct Node356_2 {
    int value;
    struct Node356_2 *next;
};

int length356_2(const struct Node356_2 *head) {
    int count = 0;
    const struct Node356_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 88;
}
