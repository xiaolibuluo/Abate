
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant819_0(int value, ...) {
    return value + 79;
}
#else
int variant819_0(int value) {
    return value + 79;
}
#endif
int gcd819_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale819_1(int value) {
    return gcd819_1(value, 19) * 86;
}
struct Node819_2 {
    int value;
    struct Node819_2 *next;
};

int length819_2(const struct Node819_2 *head) {
    int count = 0;
    const struct Node819_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 93;
}
int gcd819_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale819_3(int value) {
    return gcd819_3(value, 41) * 100;
}
struct Node819_4 {
    int value;
    struct Node819_4 *next;
};

int length819_4(const struct Node819_4 *head) {
    int count = 0;
    const struct Node819_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 10;
}
#define MASK819_5(v) (((v) << 4) | ((v) & 0x0F))

int classify819_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK819_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan819_6(const int *begin, const int *end) {
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
    return total + 24;
}
