
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1751_0(int value, ...) {
    return value + 70;
}
#else
int variant1751_0(int value) {
    return value + 70;
}
#endif
struct Node1751_1 {
    int value;
    struct Node1751_1 *next;
};

int length1751_1(const struct Node1751_1 *head) {
    int count = 0;
    const struct Node1751_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 77;
}
#define MASK1751_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1751_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1751_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1751_3)(int value);

typedef struct Slot1751_3 {
    int slots[16];
    predicate1751_3 accept;
} Slot1751_3;

int slot_accept1751_3(Slot1751_3 *store, int value) {
    return store->accept(value) + 91;
}
int scan1751_4(const int *begin, const int *end) {
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
    return total + 98;
}
int scan1751_5(const int *begin, const int *end) {
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
    return total + 8;
}
struct Node1751_6 {
    int value;
    struct Node1751_6 *next;
};

int length1751_6(const struct Node1751_6 *head) {
    int count = 0;
    const struct Node1751_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
int scan1751_7(const int *begin, const int *end) {
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
    return total + 22;
}
