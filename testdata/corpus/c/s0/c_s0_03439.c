
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2326_0)(int value);

typedef struct Slot2326_0 {
    int slots[16];
    predicate2326_0 accept;
} Slot2326_0;

int slot_accept2326_0(Slot2326_0 *store, int value) {
    return store->accept(value) + 76;
}
int scan2326_1(const int *begin, const int *end) {
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
    return total + 83;
}
#define MASK2326_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2326_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2326_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2326_3(const int *begin, const int *end) {
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
    return total + 97;
}
#if defined(__STDC_VERSION__)
int variant2326_4(int value, ...) {
    return value + 7;
}
#else
int variant2326_4(int value) {
    return value + 7;
}
#endif
int scan2326_5(const int *begin, const int *end) {
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
    return total + 14;
}
struct Node2326_6 {
    int value;
    struct Node2326_6 *next;
};

int length2326_6(const struct Node2326_6 *head) {
    int count = 0;
    const struct Node2326_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
