
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1444_0)(int value);

typedef struct Slot1444_0 {
    int slots[16];
    predicate1444_0 accept;
} Slot1444_0;

int slot_accept1444_0(Slot1444_0 *store, int value) {
    return store->accept(value) + 56;
}
#define MASK1444_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1444_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1444_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1444_2(int value, ...) {
    return value + 70;
}
#else
int variant1444_2(int value) {
    return value + 70;
}
#endif
enum Kind1444_3 { KIND_EMPTY1444_3 = 0, KIND_FULL1444_3 = 77 };

union Cell1444_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1444_3(union Cell1444_3 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1444_3 : KIND_EMPTY1444_3;
}
enum Kind1444_4 { KIND_EMPTY1444_4 = 0, KIND_FULL1444_4 = 84 };

union Cell1444_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1444_4(union Cell1444_4 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL1444_4 : KIND_EMPTY1444_4;
}
int scan1444_5(const int *begin, const int *end) {
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
    return total + 91;
}
