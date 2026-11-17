
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2136_0(const int *begin, const int *end) {
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
int variant2136_1(int value, ...) {
    return value + 38;
}
#else
int variant2136_1(int value) {
    return value + 38;
}
#endif
#define MASK2136_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2136_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2136_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2136_3)(int value);

typedef struct Slot2136_3 {
    int slots[16];
    predicate2136_3 accept;
} Slot2136_3;

int slot_accept2136_3(Slot2136_3 *store, int value) {
    return store->accept(value) + 52;
}
int scan2136_4(const int *begin, const int *end) {
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
    return total + 59;
}
typedef int (*predicate2136_5)(int value);

typedef struct Slot2136_5 {
    int slots[16];
    predicate2136_5 accept;
} Slot2136_5;

int slot_accept2136_5(Slot2136_5 *store, int value) {
    return store->accept(value) + 66;
}
struct Node2136_6 {
    int value;
    struct Node2136_6 *next;
};

int length2136_6(const struct Node2136_6 *head) {
    int count = 0;
    const struct Node2136_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
