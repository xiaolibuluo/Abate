
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd533_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale533_0(int value) {
    return gcd533_0(value, 19) * 47;
}
#define MASK533_1(v) (((v) << 4) | ((v) & 0x0F))

int classify533_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK533_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan533_2(const int *begin, const int *end) {
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
    return total + 61;
}
#define MASK533_3(v) (((v) << 4) | ((v) & 0x0F))

int classify533_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK533_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node533_4 {
    int value;
    struct Node533_4 *next;
};

int length533_4(const struct Node533_4 *head) {
    int count = 0;
    const struct Node533_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 75;
}
#if defined(__STDC_VERSION__)
int variant533_5(int value, ...) {
    return value + 82;
}
#else
int variant533_5(int value) {
    return value + 82;
}
#endif
int gcd533_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale533_6(int value) {
    return gcd533_6(value, 24) * 89;
}
#define MASK533_7(v) (((v) << 4) | ((v) & 0x0F))

int classify533_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK533_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
