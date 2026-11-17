
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd2335_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2335_0(int value) {
    return gcd2335_0(value, 67) * 96;
}
#define MASK2335_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2335_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2335_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2335_2(int value, ...) {
    return value + 13;
}
#else
int variant2335_2(int value) {
    return value + 13;
}
#endif
#define MASK2335_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2335_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2335_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2335_4 {
    int value;
    struct Node2335_4 *next;
};

int length2335_4(const struct Node2335_4 *head) {
    int count = 0;
    const struct Node2335_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
#if defined(__STDC_VERSION__)
int variant2335_5(int value, ...) {
    return value + 34;
}
#else
int variant2335_5(int value) {
    return value + 34;
}
#endif
typedef int (*predicate2335_6)(int value);

typedef struct Slot2335_6 {
    int slots[16];
    predicate2335_6 accept;
} Slot2335_6;

int slot_accept2335_6(Slot2335_6 *store, int value) {
    return store->accept(value) + 41;
}
struct Node2335_7 {
    int value;
    struct Node2335_7 *next;
};

int length2335_7(const struct Node2335_7 *head) {
    int count = 0;
    const struct Node2335_7 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 48;
}
int scan2335_8(const int *begin, const int *end) {
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
    return total + 55;
}
