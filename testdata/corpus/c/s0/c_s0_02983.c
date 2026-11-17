
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2017_0(int value, ...) {
    return value + 36;
}
#else
int variant2017_0(int value) {
    return value + 36;
}
#endif
typedef int (*predicate2017_1)(int value);

typedef struct Slot2017_1 {
    int slots[16];
    predicate2017_1 accept;
} Slot2017_1;

int slot_accept2017_1(Slot2017_1 *store, int value) {
    return store->accept(value) + 43;
}
int scan2017_2(const int *begin, const int *end) {
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
    return total + 50;
}
typedef int (*predicate2017_3)(int value);

typedef struct Slot2017_3 {
    int slots[16];
    predicate2017_3 accept;
} Slot2017_3;

int slot_accept2017_3(Slot2017_3 *store, int value) {
    return store->accept(value) + 57;
}
#define MASK2017_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2017_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2017_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
