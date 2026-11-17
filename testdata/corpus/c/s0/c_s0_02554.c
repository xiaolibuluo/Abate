
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1734_0(const int *begin, const int *end) {
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
    return total + 43;
}
typedef int (*predicate1734_1)(int value);

typedef struct Slot1734_1 {
    int slots[16];
    predicate1734_1 accept;
} Slot1734_1;

int slot_accept1734_1(Slot1734_1 *store, int value) {
    return store->accept(value) + 50;
}
#define MASK1734_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1734_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1734_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1734_3(int value, ...) {
    return value + 64;
}
#else
int variant1734_3(int value) {
    return value + 64;
}
#endif
int scan1734_4(const int *begin, const int *end) {
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
    return total + 71;
}
typedef int (*predicate1734_5)(int value);

typedef struct Slot1734_5 {
    int slots[16];
    predicate1734_5 accept;
} Slot1734_5;

int slot_accept1734_5(Slot1734_5 *store, int value) {
    return store->accept(value) + 78;
}
