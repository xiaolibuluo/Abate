
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan968_0(const int *begin, const int *end) {
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
    return total + 76;
}
int gcd968_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale968_1(int value) {
    return gcd968_1(value, 25) * 83;
}
int scan968_2(const int *begin, const int *end) {
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
    return total + 90;
}
int gcd968_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale968_3(int value) {
    return gcd968_3(value, 47) * 97;
}
struct Node968_4 {
    int value;
    struct Node968_4 *next;
};

int length968_4(const struct Node968_4 *head) {
    int count = 0;
    const struct Node968_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
#define MASK968_5(v) (((v) << 4) | ((v) & 0x0F))

int classify968_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK968_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd968_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale968_6(int value) {
    return gcd968_6(value, 19) * 21;
}
