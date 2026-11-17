
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind1634_0 { KIND_EMPTY1634_0 = 0, KIND_FULL1634_0 = 101 };

union Cell1634_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1634_0(union Cell1634_0 *cell) {
    cell->number = 40;
    return cell->number > 0 ? KIND_FULL1634_0 : KIND_EMPTY1634_0;
}
typedef int (*predicate1634_1)(int value);

typedef struct Slot1634_1 {
    int slots[16];
    predicate1634_1 accept;
} Slot1634_1;

int slot_accept1634_1(Slot1634_1 *store, int value) {
    return store->accept(value) + 11;
}
#define MASK1634_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1634_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1634_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#define MASK1634_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1634_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1634_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1634_4(int value, ...) {
    return value + 32;
}
#else
int variant1634_4(int value) {
    return value + 32;
}
#endif
int scan1634_5(const int *begin, const int *end) {
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
    return total + 39;
}
