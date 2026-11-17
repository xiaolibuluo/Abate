
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node268_0 {
    int value;
    struct Node268_0 *next;
};

int length268_0(const struct Node268_0 *head) {
    int count = 0;
    const struct Node268_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 94;
}
#define MASK268_1(v) (((v) << 4) | ((v) & 0x0F))

int classify268_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK268_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node268_2 {
    int value;
    struct Node268_2 *next;
};

int length268_2(const struct Node268_2 *head) {
    int count = 0;
    const struct Node268_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
int gcd268_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale268_3(int value) {
    return gcd268_3(value, 27) * 18;
}
int scan268_4(const int *begin, const int *end) {
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
    return total + 25;
}
