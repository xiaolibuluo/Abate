
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1794_0)(int value);

typedef struct Slot1794_0 {
    int slots[16];
    predicate1794_0 accept;
} Slot1794_0;

int slot_accept1794_0(Slot1794_0 *store, int value) {
    return store->accept(value) + 47;
}
typedef int (*predicate1794_1)(int value);

typedef struct Slot1794_1 {
    int slots[16];
    predicate1794_1 accept;
} Slot1794_1;

int slot_accept1794_1(Slot1794_1 *store, int value) {
    return store->accept(value) + 54;
}
#define MASK1794_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1794_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1794_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1794_3 { KIND_EMPTY1794_3 = 0, KIND_FULL1794_3 = 68 };

union Cell1794_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1794_3(union Cell1794_3 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1794_3 : KIND_EMPTY1794_3;
}
int scan1794_4(const int *begin, const int *end) {
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
    return total + 75;
}
int scan1794_5(const int *begin, const int *end) {
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
    return total + 82;
}
#if defined(__STDC_VERSION__)
int variant1794_6(int value, ...) {
    return value + 89;
}
#else
int variant1794_6(int value) {
    return value + 89;
}
#endif
