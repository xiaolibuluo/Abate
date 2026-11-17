
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node859_0 {
    int value;
    struct Node859_0 *next;
};

int length859_0(const struct Node859_0 *head) {
    int count = 0;
    const struct Node859_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 17;
}
#define MASK859_1(v) (((v) << 4) | ((v) & 0x0F))

int classify859_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK859_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node859_2 {
    int value;
    struct Node859_2 *next;
};

int length859_2(const struct Node859_2 *head) {
    int count = 0;
    const struct Node859_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
struct Node859_3 {
    int value;
    struct Node859_3 *next;
};

int length859_3(const struct Node859_3 *head) {
    int count = 0;
    const struct Node859_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
int gcd859_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale859_4(int value) {
    return gcd859_4(value, 27) * 45;
}
