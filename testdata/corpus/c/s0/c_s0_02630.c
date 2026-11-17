
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1784_0(int value, ...) {
    return value + 14;
}
#else
int variant1784_0(int value) {
    return value + 14;
}
#endif
typedef int (*predicate1784_1)(int value);

typedef struct Slot1784_1 {
    int slots[16];
    predicate1784_1 accept;
} Slot1784_1;

int slot_accept1784_1(Slot1784_1 *store, int value) {
    return store->accept(value) + 21;
}
enum Kind1784_2 { KIND_EMPTY1784_2 = 0, KIND_FULL1784_2 = 28 };

union Cell1784_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1784_2(union Cell1784_2 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL1784_2 : KIND_EMPTY1784_2;
}
typedef int (*predicate1784_3)(int value);

typedef struct Slot1784_3 {
    int slots[16];
    predicate1784_3 accept;
} Slot1784_3;

int slot_accept1784_3(Slot1784_3 *store, int value) {
    return store->accept(value) + 35;
}
#define MASK1784_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1784_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1784_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1784_5(const int *begin, const int *end) {
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
    return total + 49;
}
