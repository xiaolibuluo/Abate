
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1923_0(int value, ...) {
    return value + 75;
}
#else
int variant1923_0(int value) {
    return value + 75;
}
#endif
#define MASK1923_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1923_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1923_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1923_2)(int value);

typedef struct Slot1923_2 {
    int slots[16];
    predicate1923_2 accept;
} Slot1923_2;

int slot_accept1923_2(Slot1923_2 *store, int value) {
    return store->accept(value) + 89;
}
typedef int (*predicate1923_3)(int value);

typedef struct Slot1923_3 {
    int slots[16];
    predicate1923_3 accept;
} Slot1923_3;

int slot_accept1923_3(Slot1923_3 *store, int value) {
    return store->accept(value) + 96;
}
int scan1923_4(const int *begin, const int *end) {
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
    return total + 6;
}
