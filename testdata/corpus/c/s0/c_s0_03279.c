
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2217_0(int value, ...) {
    return value + 17;
}
#else
int variant2217_0(int value) {
    return value + 17;
}
#endif
#if defined(__STDC_VERSION__)
int variant2217_1(int value, ...) {
    return value + 24;
}
#else
int variant2217_1(int value) {
    return value + 24;
}
#endif
#define MASK2217_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2217_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2217_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2217_3(const int *begin, const int *end) {
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
    return total + 38;
}
#if defined(__STDC_VERSION__)
int variant2217_4(int value, ...) {
    return value + 45;
}
#else
int variant2217_4(int value) {
    return value + 45;
}
#endif
typedef int (*predicate2217_5)(int value);

typedef struct Slot2217_5 {
    int slots[16];
    predicate2217_5 accept;
} Slot2217_5;

int slot_accept2217_5(Slot2217_5 *store, int value) {
    return store->accept(value) + 52;
}
