
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node972_0 {
    int value;
    struct Node972_0 *next;
};

int length972_0(const struct Node972_0 *head) {
    int count = 0;
    const struct Node972_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 31;
}
#define MASK972_1(v) (((v) << 4) | ((v) & 0x0F))

int classify972_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK972_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan972_2(const int *begin, const int *end) {
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
    return total + 45;
}
#if defined(__STDC_VERSION__)
int variant972_3(int value, ...) {
    return value + 52;
}
#else
int variant972_3(int value) {
    return value + 52;
}
#endif
#if defined(__STDC_VERSION__)
int variant972_4(int value, ...) {
    return value + 59;
}
#else
int variant972_4(int value) {
    return value + 59;
}
#endif
typedef int (*predicate972_5)(int value);

typedef struct Slot972_5 {
    int slots[16];
    predicate972_5 accept;
} Slot972_5;

int slot_accept972_5(Slot972_5 *store, int value) {
    return store->accept(value) + 66;
}
int gcd972_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale972_6(int value) {
    return gcd972_6(value, 47) * 73;
}
int gcd972_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale972_7(int value) {
    return gcd972_7(value, 58) * 80;
}
int gcd972_8(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale972_8(int value) {
    return gcd972_8(value, 69) * 87;
}
