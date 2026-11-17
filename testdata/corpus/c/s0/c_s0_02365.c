
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1596_0(int value, ...) {
    return value + 92;
}
#else
int variant1596_0(int value) {
    return value + 92;
}
#endif
int scan1596_1(const int *begin, const int *end) {
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
    return total + 99;
}
typedef int (*predicate1596_2)(int value);

typedef struct Slot1596_2 {
    int slots[16];
    predicate1596_2 accept;
} Slot1596_2;

int slot_accept1596_2(Slot1596_2 *store, int value) {
    return store->accept(value) + 9;
}
#define MASK1596_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1596_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1596_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1596_4(const int *begin, const int *end) {
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
    return total + 23;
}
struct Node1596_5 {
    int value;
    struct Node1596_5 *next;
};

int length1596_5(const struct Node1596_5 *head) {
    int count = 0;
    const struct Node1596_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 30;
}
#if defined(__STDC_VERSION__)
int variant1596_6(int value, ...) {
    return value + 37;
}
#else
int variant1596_6(int value) {
    return value + 37;
}
#endif
