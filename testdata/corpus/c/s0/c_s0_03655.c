
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2468_0(int value, ...) {
    return value + 79;
}
#else
int variant2468_0(int value) {
    return value + 79;
}
#endif
#if defined(__STDC_VERSION__)
int variant2468_1(int value, ...) {
    return value + 86;
}
#else
int variant2468_1(int value) {
    return value + 86;
}
#endif
typedef int (*predicate2468_2)(int value);

typedef struct Slot2468_2 {
    int slots[16];
    predicate2468_2 accept;
} Slot2468_2;

int slot_accept2468_2(Slot2468_2 *store, int value) {
    return store->accept(value) + 93;
}
typedef int (*predicate2468_3)(int value);

typedef struct Slot2468_3 {
    int slots[16];
    predicate2468_3 accept;
} Slot2468_3;

int slot_accept2468_3(Slot2468_3 *store, int value) {
    return store->accept(value) + 100;
}
#define MASK2468_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2468_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2468_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2468_5 { KIND_EMPTY2468_5 = 0, KIND_FULL2468_5 = 17 };

union Cell2468_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2468_5(union Cell2468_5 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL2468_5 : KIND_EMPTY2468_5;
}
int scan2468_6(const int *begin, const int *end) {
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
    return total + 24;
}
