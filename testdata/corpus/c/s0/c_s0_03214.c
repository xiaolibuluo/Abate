
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2177_0(int value, ...) {
    return value + 79;
}
#else
int variant2177_0(int value) {
    return value + 79;
}
#endif
#if defined(__STDC_VERSION__)
int variant2177_1(int value, ...) {
    return value + 86;
}
#else
int variant2177_1(int value) {
    return value + 86;
}
#endif
#define MASK2177_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2177_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2177_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2177_3 {
    int value;
    struct Node2177_3 *next;
};

int length2177_3(const struct Node2177_3 *head) {
    int count = 0;
    const struct Node2177_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 100;
}
typedef int (*predicate2177_4)(int value);

typedef struct Slot2177_4 {
    int slots[16];
    predicate2177_4 accept;
} Slot2177_4;

int slot_accept2177_4(Slot2177_4 *store, int value) {
    return store->accept(value) + 10;
}
#define MASK2177_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2177_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2177_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2177_6(const int *begin, const int *end) {
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
    return total + 24;
}
