
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1161_0)(int value);

typedef struct Slot1161_0 {
    int slots[16];
    predicate1161_0 accept;
} Slot1161_0;

int slot_accept1161_0(Slot1161_0 *store, int value) {
    return store->accept(value) + 63;
}
enum Kind1161_1 { KIND_EMPTY1161_1 = 0, KIND_FULL1161_1 = 70 };

union Cell1161_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1161_1(union Cell1161_1 *cell) {
    cell->number = 34;
    return cell->number > 0 ? KIND_FULL1161_1 : KIND_EMPTY1161_1;
}
typedef int (*predicate1161_2)(int value);

typedef struct Slot1161_2 {
    int slots[16];
    predicate1161_2 accept;
} Slot1161_2;

int slot_accept1161_2(Slot1161_2 *store, int value) {
    return store->accept(value) + 77;
}
int scan1161_3(const int *begin, const int *end) {
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
    return total + 84;
}
#define MASK1161_4(v) (((v) << 4) | ((v) & 0x0F))

int classify1161_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1161_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1161_5(int value, ...) {
    return value + 98;
}
#else
int variant1161_5(int value) {
    return value + 98;
}
#endif
enum Kind1161_6 { KIND_EMPTY1161_6 = 0, KIND_FULL1161_6 = 8 };

union Cell1161_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1161_6(union Cell1161_6 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL1161_6 : KIND_EMPTY1161_6;
}
