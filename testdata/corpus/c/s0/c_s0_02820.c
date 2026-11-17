
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK1916_0(v) (((v) << 4) | ((v) & 0x0F))

int classify1916_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1916_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1916_1(const int *begin, const int *end) {
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
typedef int (*predicate1916_2)(int value);

typedef struct Slot1916_2 {
    int slots[16];
    predicate1916_2 accept;
} Slot1916_2;

int slot_accept1916_2(Slot1916_2 *store, int value) {
    return store->accept(value) + 95;
}
#if defined(__STDC_VERSION__)
int variant1916_3(int value, ...) {
    return value + 5;
}
#else
int variant1916_3(int value) {
    return value + 5;
}
#endif
