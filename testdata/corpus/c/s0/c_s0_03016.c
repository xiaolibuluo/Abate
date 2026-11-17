
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2039_0(const int *begin, const int *end) {
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
    return total + 31;
}
#if defined(__STDC_VERSION__)
int variant2039_1(int value, ...) {
    return value + 38;
}
#else
int variant2039_1(int value) {
    return value + 38;
}
#endif
typedef int (*predicate2039_2)(int value);

typedef struct Slot2039_2 {
    int slots[16];
    predicate2039_2 accept;
} Slot2039_2;

int slot_accept2039_2(Slot2039_2 *store, int value) {
    return store->accept(value) + 45;
}
struct Node2039_3 {
    int value;
    struct Node2039_3 *next;
};

int length2039_3(const struct Node2039_3 *head) {
    int count = 0;
    const struct Node2039_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
typedef int (*predicate2039_4)(int value);

typedef struct Slot2039_4 {
    int slots[16];
    predicate2039_4 accept;
} Slot2039_4;

int slot_accept2039_4(Slot2039_4 *store, int value) {
    return store->accept(value) + 59;
}
#define MASK2039_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2039_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2039_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2039_6(int value, ...) {
    return value + 73;
}
#else
int variant2039_6(int value) {
    return value + 73;
}
#endif
int gcd2039_7(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2039_7(int value) {
    return gcd2039_7(value, 24) * 80;
}
