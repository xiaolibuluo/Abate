
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd993_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale993_0(int value) {
    return gcd993_0(value, 67) * 13;
}
#define MASK993_1(v) (((v) << 4) | ((v) & 0x0F))

int classify993_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK993_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node993_2 {
    int value;
    struct Node993_2 *next;
};

int length993_2(const struct Node993_2 *head) {
    int count = 0;
    const struct Node993_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
int gcd993_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale993_3(int value) {
    return gcd993_3(value, 39) * 34;
}
