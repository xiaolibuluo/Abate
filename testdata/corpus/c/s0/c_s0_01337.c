
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan915_0(const int *begin, const int *end) {
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
    return total + 66;
}
#define MASK915_1(v) (((v) << 4) | ((v) & 0x0F))

int classify915_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK915_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK915_2(v) (((v) << 4) | ((v) & 0x0F))

int classify915_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK915_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK915_3(v) (((v) << 4) | ((v) & 0x0F))

int classify915_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK915_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd915_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale915_4(int value) {
    return gcd915_4(value, 53) * 94;
}
int scan915_5(const int *begin, const int *end) {
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
    return total + 101;
}
struct Node915_6 {
    int value;
    struct Node915_6 *next;
};

int length915_6(const struct Node915_6 *head) {
    int count = 0;
    const struct Node915_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
