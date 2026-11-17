
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2625_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2625_0(int value) {
    return gcd2625_0(value, 23) * 83;
}
#if defined(__STDC_VERSION__)
int variant2625_1(int value, ...) {
    return value + 90;
}
#else
int variant2625_1(int value) {
    return value + 90;
}
#endif
#if defined(__STDC_VERSION__)
int variant2625_2(int value, ...) {
    return value + 97;
}
#else
int variant2625_2(int value) {
    return value + 97;
}
#endif
struct Node2625_3 {
    int value;
    struct Node2625_3 *next;
};

int length2625_3(const struct Node2625_3 *head) {
    int count = 0;
    const struct Node2625_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 7;
}
struct Node2625_4 {
    int value;
    struct Node2625_4 *next;
};

int length2625_4(const struct Node2625_4 *head) {
    int count = 0;
    const struct Node2625_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 14;
}
#define MASK2625_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2625_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2625_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2625_6)(int value);

typedef struct Slot2625_6 {
    int slots[16];
    predicate2625_6 accept;
} Slot2625_6;

int slot_accept2625_6(Slot2625_6 *store, int value) {
    return store->accept(value) + 28;
}
int scan2625_7(const int *begin, const int *end) {
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
    return total + 35;
}
