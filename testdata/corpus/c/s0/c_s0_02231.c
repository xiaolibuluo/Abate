
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1500_0(int value, ...) {
    return value + 8;
}
#else
int variant1500_0(int value) {
    return value + 8;
}
#endif
#if defined(__STDC_VERSION__)
int variant1500_1(int value, ...) {
    return value + 15;
}
#else
int variant1500_1(int value) {
    return value + 15;
}
#endif
typedef int (*predicate1500_2)(int value);

typedef struct Slot1500_2 {
    int slots[16];
    predicate1500_2 accept;
} Slot1500_2;

int slot_accept1500_2(Slot1500_2 *store, int value) {
    return store->accept(value) + 22;
}
#define MASK1500_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1500_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1500_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1500_4)(int value);

typedef struct Slot1500_4 {
    int slots[16];
    predicate1500_4 accept;
} Slot1500_4;

int slot_accept1500_4(Slot1500_4 *store, int value) {
    return store->accept(value) + 36;
}
typedef int (*predicate1500_5)(int value);

typedef struct Slot1500_5 {
    int slots[16];
    predicate1500_5 accept;
} Slot1500_5;

int slot_accept1500_5(Slot1500_5 *store, int value) {
    return store->accept(value) + 43;
}
#if defined(__STDC_VERSION__)
int variant1500_6(int value, ...) {
    return value + 50;
}
#else
int variant1500_6(int value) {
    return value + 50;
}
#endif
int scan1500_7(const int *begin, const int *end) {
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
    return total + 57;
}
