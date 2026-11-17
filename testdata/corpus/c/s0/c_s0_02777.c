
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1890_0(const int *begin, const int *end) {
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
    return total + 34;
}
#if defined(__STDC_VERSION__)
int variant1890_1(int value, ...) {
    return value + 41;
}
#else
int variant1890_1(int value) {
    return value + 41;
}
#endif
int scan1890_2(const int *begin, const int *end) {
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
    return total + 48;
}
#if defined(__STDC_VERSION__)
int variant1890_3(int value, ...) {
    return value + 55;
}
#else
int variant1890_3(int value) {
    return value + 55;
}
#endif
int scan1890_4(const int *begin, const int *end) {
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
    return total + 62;
}
#define MASK1890_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1890_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1890_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1890_6)(int value);

typedef struct Slot1890_6 {
    int slots[16];
    predicate1890_6 accept;
} Slot1890_6;

int slot_accept1890_6(Slot1890_6 *store, int value) {
    return store->accept(value) + 76;
}
