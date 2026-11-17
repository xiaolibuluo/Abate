
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node770_0 {
    int value;
    struct Node770_0 *next;
};

int length770_0(const struct Node770_0 *head) {
    int count = 0;
    const struct Node770_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 24;
}
int scan770_1(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 31;
}
struct Node770_2 {
    int value;
    struct Node770_2 *next;
};

int length770_2(const struct Node770_2 *head) {
    int count = 0;
    const struct Node770_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
int gcd770_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale770_3(int value) {
    return gcd770_3(value, 64) * 45;
}
#define MASK770_4(v) (((v) << 4) | ((v) & 0x0F))

int classify770_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK770_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd770_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale770_5(int value) {
    return gcd770_5(value, 25) * 59;
}
struct Node770_6 {
    int value;
    struct Node770_6 *next;
};

int length770_6(const struct Node770_6 *head) {
    int count = 0;
    const struct Node770_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 66;
}
