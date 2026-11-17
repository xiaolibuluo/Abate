
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1696_0)(int value);

typedef struct Slot1696_0 {
    int slots[16];
    predicate1696_0 accept;
} Slot1696_0;

int slot_accept1696_0(Slot1696_0 *store, int value) {
    return store->accept(value) + 34;
}
#if defined(__STDC_VERSION__)
int variant1696_1(int value, ...) {
    return value + 41;
}
#else
int variant1696_1(int value) {
    return value + 41;
}
#endif
#define MASK1696_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1696_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1696_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
struct Node1696_3 {
    int value;
    struct Node1696_3 *next;
};

int length1696_3(const struct Node1696_3 *head) {
    int count = 0;
    const struct Node1696_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 55;
}
#define MASK1696_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1696_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1696_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1696_5(const int *begin, const int *end) {
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
    return total + 69;
}
int scan1696_6(const int *begin, const int *end) {
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
    return total + 76;
}
