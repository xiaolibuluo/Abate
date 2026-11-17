
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK817_0(v) (((v) << 4) | ((v) & 0x0F))

int classify817_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK817_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK817_1(v) (((v) << 4) | ((v) & 0x0F))

int classify817_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK817_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate817_2)(int value);

typedef struct Slot817_2 {
    int slots[16];
    predicate817_2 accept;
} Slot817_2;

int slot_accept817_2(Slot817_2 *store, int value) {
    return store->accept(value) + 67;
}
typedef int (*predicate817_3)(int value);

typedef struct Slot817_3 {
    int slots[16];
    predicate817_3 accept;
} Slot817_3;

int slot_accept817_3(Slot817_3 *store, int value) {
    return store->accept(value) + 74;
}
int scan817_4(const int *begin, const int *end) {
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
    return total + 81;
}
int scan817_5(const int *begin, const int *end) {
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
    return total + 88;
}
#if defined(__STDC_VERSION__)
int variant817_6(int value, ...) {
    return value + 95;
}
#else
int variant817_6(int value) {
    return value + 95;
}
#endif
