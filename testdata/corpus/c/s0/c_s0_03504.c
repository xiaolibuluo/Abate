
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan2364_0(const int *begin, const int *end) {
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
    return total + 85;
}
#define MASK2364_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2364_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2364_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2364_2(int value, ...) {
    return value + 99;
}
#else
int variant2364_2(int value) {
    return value + 99;
}
#endif
int scan2364_3(const int *begin, const int *end) {
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
    return total + 9;
}
typedef int (*predicate2364_4)(int value);

typedef struct Slot2364_4 {
    int slots[16];
    predicate2364_4 accept;
} Slot2364_4;

int slot_accept2364_4(Slot2364_4 *store, int value) {
    return store->accept(value) + 16;
}
