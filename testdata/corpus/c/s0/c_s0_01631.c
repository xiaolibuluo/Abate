
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1100_0 {
    int value;
    struct Node1100_0 *next;
};

int length1100_0(const struct Node1100_0 *head) {
    int count = 0;
    const struct Node1100_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 46;
}
int gcd1100_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1100_1(int value) {
    return gcd1100_1(value, 34) * 53;
}
struct Node1100_2 {
    int value;
    struct Node1100_2 *next;
};

int length1100_2(const struct Node1100_2 *head) {
    int count = 0;
    const struct Node1100_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 60;
}
#define MASK1100_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1100_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1100_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1100_4 {
    int value;
    struct Node1100_4 *next;
};

int length1100_4(const struct Node1100_4 *head) {
    int count = 0;
    const struct Node1100_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 74;
}
