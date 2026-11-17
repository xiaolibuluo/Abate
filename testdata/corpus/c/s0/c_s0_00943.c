
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd648_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale648_0(int value) {
    return gcd648_0(value, 31) * 87;
}
#define MASK648_1(v) (((v) << 4) | ((v) & 0x0F))

int classify648_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK648_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node648_2 {
    int value;
    struct Node648_2 *next;
};

int length648_2(const struct Node648_2 *head) {
    int count = 0;
    const struct Node648_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
