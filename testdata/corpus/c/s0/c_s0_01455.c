
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node989_0 {
    int value;
    struct Node989_0 *next;
};

int length989_0(const struct Node989_0 *head) {
    int count = 0;
    const struct Node989_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 58;
}
#define MASK989_1(v) (((v) << 4) | ((v) & 0x0F))

int classify989_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK989_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd989_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale989_2(int value) {
    return gcd989_2(value, 61) * 72;
}
