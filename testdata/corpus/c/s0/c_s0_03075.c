
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2076_0(const int *begin, const int *end) {
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
    return total + 27;
}
int scan2076_1(const int *begin, const int *end) {
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
    return total + 34;
}
typedef int (*predicate2076_2)(int value);

typedef struct Slot2076_2 {
    int slots[16];
    predicate2076_2 accept;
} Slot2076_2;

int slot_accept2076_2(Slot2076_2 *store, int value) {
    return store->accept(value) + 41;
}
#define MASK2076_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2076_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2076_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2076_4(int value, ...) {
    return value + 55;
}
#else
int variant2076_4(int value) {
    return value + 55;
}
#endif
struct Node2076_5 {
    int value;
    struct Node2076_5 *next;
};

int length2076_5(const struct Node2076_5 *head) {
    int count = 0;
    const struct Node2076_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 62;
}
