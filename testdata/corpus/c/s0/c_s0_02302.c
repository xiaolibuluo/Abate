
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1552_0(const int *begin, const int *end) {
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
    return total + 5;
}
#if defined(__STDC_VERSION__)
int variant1552_1(int value, ...) {
    return value + 12;
}
#else
int variant1552_1(int value) {
    return value + 12;
}
#endif
int scan1552_2(const int *begin, const int *end) {
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
    return total + 19;
}
int gcd1552_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1552_3(int value) {
    return gcd1552_3(value, 48) * 26;
}
#define MASK1552_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1552_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1552_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1552_5(int value, ...) {
    return value + 40;
}
#else
int variant1552_5(int value) {
    return value + 40;
}
#endif
struct Node1552_6 {
    int value;
    struct Node1552_6 *next;
};

int length1552_6(const struct Node1552_6 *head) {
    int count = 0;
    const struct Node1552_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 47;
}
typedef int (*predicate1552_7)(int value);

typedef struct Slot1552_7 {
    int slots[16];
    predicate1552_7 accept;
} Slot1552_7;

int slot_accept1552_7(Slot1552_7 *store, int value) {
    return store->accept(value) + 54;
}
