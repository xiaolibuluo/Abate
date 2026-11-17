
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2471_0(const int *begin, const int *end) {
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
    return total + 21;
}
#define MASK2471_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2471_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2471_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2471_2)(int value);

typedef struct Slot2471_2 {
    int slots[16];
    predicate2471_2 accept;
} Slot2471_2;

int slot_accept2471_2(Slot2471_2 *store, int value) {
    return store->accept(value) + 35;
}
struct Node2471_3 {
    int value;
    struct Node2471_3 *next;
};

int length2471_3(const struct Node2471_3 *head) {
    int count = 0;
    const struct Node2471_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 42;
}
#define MASK2471_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2471_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2471_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2471_5(int value, ...) {
    return value + 56;
}
#else
int variant2471_5(int value) {
    return value + 56;
}
#endif
#define MASK2471_6(v) (((v) << 4) | ((v) & 0x0F))

int classify2471_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2471_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
