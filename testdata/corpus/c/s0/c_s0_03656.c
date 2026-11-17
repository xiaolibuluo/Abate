
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2469_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2469_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2469_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2469_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2469_1(int value) {
    return gcd2469_1(value, 40) * 99;
}
#define MASK2469_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2469_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2469_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2469_3(int value, ...) {
    return value + 16;
}
#else
int variant2469_3(int value) {
    return value + 16;
}
#endif
struct Node2469_4 {
    int value;
    struct Node2469_4 *next;
};

int length2469_4(const struct Node2469_4 *head) {
    int count = 0;
    const struct Node2469_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
typedef int (*predicate2469_5)(int value);

typedef struct Slot2469_5 {
    int slots[16];
    predicate2469_5 accept;
} Slot2469_5;

int slot_accept2469_5(Slot2469_5 *store, int value) {
    return store->accept(value) + 30;
}
int scan2469_6(const int *begin, const int *end) {
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
    return total + 37;
}
