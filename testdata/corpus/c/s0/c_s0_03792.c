
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2567_0(int value, ...) {
    return value + 8;
}
#else
int variant2567_0(int value) {
    return value + 8;
}
#endif
int scan2567_1(const int *begin, const int *end) {
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
    return total + 15;
}
#if defined(__STDC_VERSION__)
int variant2567_2(int value, ...) {
    return value + 22;
}
#else
int variant2567_2(int value) {
    return value + 22;
}
#endif
#if defined(__STDC_VERSION__)
int variant2567_3(int value, ...) {
    return value + 29;
}
#else
int variant2567_3(int value) {
    return value + 29;
}
#endif
typedef int (*predicate2567_4)(int value);

typedef struct Slot2567_4 {
    int slots[16];
    predicate2567_4 accept;
} Slot2567_4;

int slot_accept2567_4(Slot2567_4 *store, int value) {
    return store->accept(value) + 36;
}
#define MASK2567_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2567_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2567_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node2567_6 {
    int value;
    struct Node2567_6 *next;
};

int length2567_6(const struct Node2567_6 *head) {
    int count = 0;
    const struct Node2567_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
