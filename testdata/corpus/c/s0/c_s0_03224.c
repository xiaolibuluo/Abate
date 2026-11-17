
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2184_0(int value, ...) {
    return value + 73;
}
#else
int variant2184_0(int value) {
    return value + 73;
}
#endif
int scan2184_1(const int *begin, const int *end) {
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
    return total + 80;
}
typedef int (*predicate2184_2)(int value);

typedef struct Slot2184_2 {
    int slots[16];
    predicate2184_2 accept;
} Slot2184_2;

int slot_accept2184_2(Slot2184_2 *store, int value) {
    return store->accept(value) + 87;
}
int scan2184_3(const int *begin, const int *end) {
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
    return total + 94;
}
#define MASK2184_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2184_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2184_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2184_5 {
    int value;
    struct Node2184_5 *next;
};

int length2184_5(const struct Node2184_5 *head) {
    int count = 0;
    const struct Node2184_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 11;
}
int gcd2184_6(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2184_6(int value) {
    return gcd2184_6(value, 52) * 18;
}
