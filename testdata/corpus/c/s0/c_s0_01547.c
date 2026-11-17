
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1046_0)(int value);

typedef struct Slot1046_0 {
    int slots[16];
    predicate1046_0 accept;
} Slot1046_0;

int slot_accept1046_0(Slot1046_0 *store, int value) {
    return store->accept(value) + 23;
}
#define MASK1046_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1046_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1046_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1046_2(int value, ...) {
    return value + 37;
}
#else
int variant1046_2(int value) {
    return value + 37;
}
#endif
#if defined(__STDC_VERSION__)
int variant1046_3(int value, ...) {
    return value + 44;
}
#else
int variant1046_3(int value) {
    return value + 44;
}
#endif
int scan1046_4(const int *begin, const int *end) {
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
