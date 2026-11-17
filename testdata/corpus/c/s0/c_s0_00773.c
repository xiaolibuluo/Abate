
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd535_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale535_0(int value) {
    return gcd535_0(value, 33) * 73;
}
#define MASK535_1(v) (((v) << 4) | ((v) & 0x0F))

int classify535_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK535_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd535_2(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale535_2(int value) {
    return gcd535_2(value, 55) * 87;
}
#define MASK535_3(v) (((v) << 4) | ((v) & 0x0F))

int classify535_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK535_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node535_4 {
    int value;
    struct Node535_4 *next;
};

int length535_4(const struct Node535_4 *head) {
    int count = 0;
    const struct Node535_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
int scan535_5(const int *begin, const int *end) {
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
    return total + 11;
}
struct Node535_6 {
    int value;
    struct Node535_6 *next;
};

int length535_6(const struct Node535_6 *head) {
    int count = 0;
    const struct Node535_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
int gcd535_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale535_7(int value) {
    return gcd535_7(value, 49) * 25;
}
