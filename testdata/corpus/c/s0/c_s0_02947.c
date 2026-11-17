
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1992_0(int value, ...) {
    return value + 99;
}
#else
int variant1992_0(int value) {
    return value + 99;
}
#endif
#define MASK1992_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1992_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1992_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1992_2(int value, ...) {
    return value + 16;
}
#else
int variant1992_2(int value) {
    return value + 16;
}
#endif
struct Node1992_3 {
    int value;
    struct Node1992_3 *next;
};

int length1992_3(const struct Node1992_3 *head) {
    int count = 0;
    const struct Node1992_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 23;
}
typedef int (*predicate1992_4)(int value);

typedef struct Slot1992_4 {
    int slots[16];
    predicate1992_4 accept;
} Slot1992_4;

int slot_accept1992_4(Slot1992_4 *store, int value) {
    return store->accept(value) + 30;
}
int gcd1992_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1992_5(int value) {
    return gcd1992_5(value, 39) * 37;
}
struct Node1992_6 {
    int value;
    struct Node1992_6 *next;
};

int length1992_6(const struct Node1992_6 *head) {
    int count = 0;
    const struct Node1992_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 44;
}
int scan1992_7(const int *begin, const int *end) {
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
    return total + 51;
}
