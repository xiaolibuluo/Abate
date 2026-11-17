
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node848_0 {
    int value;
    struct Node848_0 *next;
};

int length848_0(const struct Node848_0 *head) {
    int count = 0;
    const struct Node848_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
int gcd848_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale848_1(int value) {
    return gcd848_1(value, 39) * 75;
}
#define MASK848_2(v) (((v) << 4) | ((v) & 0x0F))

int classify848_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK848_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd848_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale848_3(int value) {
    return gcd848_3(value, 61) * 89;
}
