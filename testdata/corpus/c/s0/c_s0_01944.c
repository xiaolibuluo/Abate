
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1313_0)(int value);

typedef struct Slot1313_0 {
    int slots[16];
    predicate1313_0 accept;
} Slot1313_0;

int slot_accept1313_0(Slot1313_0 *store, int value) {
    return store->accept(value) + 99;
}
typedef int (*predicate1313_1)(int value);

typedef struct Slot1313_1 {
    int slots[16];
    predicate1313_1 accept;
} Slot1313_1;

int slot_accept1313_1(Slot1313_1 *store, int value) {
    return store->accept(value) + 9;
}
#if defined(__STDC_VERSION__)
int variant1313_2(int value, ...) {
    return value + 16;
}
#else
int variant1313_2(int value) {
    return value + 16;
}
#endif
int scan1313_3(const int *begin, const int *end) {
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
#define MASK1313_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1313_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1313_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1313_5(int value, ...) {
    return value + 37;
}
#else
int variant1313_5(int value) {
    return value + 37;
}
#endif
#if defined(__STDC_VERSION__)
int variant1313_6(int value, ...) {
    return value + 44;
}
#else
int variant1313_6(int value) {
    return value + 44;
}
#endif
int scan1313_7(const int *begin, const int *end) {
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
    return total + 51;
}
#if defined(__STDC_VERSION__)
int variant1313_8(int value, ...) {
    return value + 58;
}
#else
int variant1313_8(int value) {
    return value + 58;
}
#endif
