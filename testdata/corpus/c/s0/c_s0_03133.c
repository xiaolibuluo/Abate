
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2123_0(const int *begin, const int *end) {
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
    return total + 56;
}
#if defined(__STDC_VERSION__)
int variant2123_1(int value, ...) {
    return value + 63;
}
#else
int variant2123_1(int value) {
    return value + 63;
}
#endif
int scan2123_2(const int *begin, const int *end) {
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
    return total + 70;
}
#define MASK2123_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2123_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2123_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2123_4)(int value);

typedef struct Slot2123_4 {
    int slots[16];
    predicate2123_4 accept;
} Slot2123_4;

int slot_accept2123_4(Slot2123_4 *store, int value) {
    return store->accept(value) + 84;
}
#define MASK2123_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2123_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2123_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
