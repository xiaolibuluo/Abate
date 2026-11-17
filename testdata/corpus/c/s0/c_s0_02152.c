
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1448_0(int value, ...) {
    return value + 11;
}
#else
int variant1448_0(int value) {
    return value + 11;
}
#endif
#define MASK1448_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1448_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1448_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1448_2)(int value);

typedef struct Slot1448_2 {
    int slots[16];
    predicate1448_2 accept;
} Slot1448_2;

int slot_accept1448_2(Slot1448_2 *store, int value) {
    return store->accept(value) + 25;
}
struct Node1448_3 {
    int value;
    struct Node1448_3 *next;
};

int length1448_3(const struct Node1448_3 *head) {
    int count = 0;
    const struct Node1448_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 32;
}
int scan1448_4(const int *begin, const int *end) {
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
    return total + 39;
}
typedef int (*predicate1448_5)(int value);

typedef struct Slot1448_5 {
    int slots[16];
    predicate1448_5 accept;
} Slot1448_5;

int slot_accept1448_5(Slot1448_5 *store, int value) {
    return store->accept(value) + 46;
}
#if defined(__STDC_VERSION__)
int variant1448_6(int value, ...) {
    return value + 53;
}
#else
int variant1448_6(int value) {
    return value + 53;
}
#endif
