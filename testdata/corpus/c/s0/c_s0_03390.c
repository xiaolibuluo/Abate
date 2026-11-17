
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2286_0(int value, ...) {
    return value + 41;
}
#else
int variant2286_0(int value) {
    return value + 41;
}
#endif
#if defined(__STDC_VERSION__)
int variant2286_1(int value, ...) {
    return value + 48;
}
#else
int variant2286_1(int value) {
    return value + 48;
}
#endif
int scan2286_2(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant2286_3(int value, ...) {
    return value + 62;
}
#else
int variant2286_3(int value) {
    return value + 62;
}
#endif
#define MASK2286_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2286_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2286_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int gcd2286_5(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2286_5(int value) {
    return gcd2286_5(value, 23) * 76;
}
struct Node2286_6 {
    int value;
    struct Node2286_6 *next;
};

int length2286_6(const struct Node2286_6 *head) {
    int count = 0;
    const struct Node2286_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 83;
}
#define MASK2286_7(v) (((v) << 4) | ((v) & 0x0F))

int classify2286_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2286_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2286_8)(int value);

typedef struct Slot2286_8 {
    int slots[16];
    predicate2286_8 accept;
} Slot2286_8;

int slot_accept2286_8(Slot2286_8 *store, int value) {
    return store->accept(value) + 97;
}
