
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1791_0(const int *begin, const int *end) {
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
    return total + 8;
}
#if defined(__STDC_VERSION__)
int variant1791_1(int value, ...) {
    return value + 15;
}
#else
int variant1791_1(int value) {
    return value + 15;
}
#endif
enum Kind1791_2 { KIND_EMPTY1791_2 = 0, KIND_FULL1791_2 = 22 };

union Cell1791_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1791_2(union Cell1791_2 *cell) {
    cell->number = 63;
    return cell->number > 0 ? KIND_FULL1791_2 : KIND_EMPTY1791_2;
}
enum Kind1791_3 { KIND_EMPTY1791_3 = 0, KIND_FULL1791_3 = 29 };

union Cell1791_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1791_3(union Cell1791_3 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1791_3 : KIND_EMPTY1791_3;
}
enum Kind1791_4 { KIND_EMPTY1791_4 = 0, KIND_FULL1791_4 = 36 };

union Cell1791_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1791_4(union Cell1791_4 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL1791_4 : KIND_EMPTY1791_4;
}
#define MASK1791_5(v) (((v) << 4) | ((v) & 0x0F))

int classify1791_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1791_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate1791_6)(int value);

typedef struct Slot1791_6 {
    int slots[16];
    predicate1791_6 accept;
} Slot1791_6;

int slot_accept1791_6(Slot1791_6 *store, int value) {
    return store->accept(value) + 50;
}
