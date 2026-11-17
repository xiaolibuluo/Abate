
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan953_0(const int *begin, const int *end) {
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
    return total + 75;
}
#define MASK953_1(v) (((v) << 4) | ((v) & 0x0F))

int classify953_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK953_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node953_2 {
    int value;
    struct Node953_2 *next;
};

int length953_2(const struct Node953_2 *head) {
    int count = 0;
    const struct Node953_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 89;
}
int gcd953_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale953_3(int value) {
    return gcd953_3(value, 64) * 96;
}
int gcd953_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale953_4(int value) {
    return gcd953_4(value, 14) * 6;
}
