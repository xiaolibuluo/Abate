
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan1700_0(const int *begin, const int *end) {
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
    return total + 86;
}
#define MASK1700_1(v) (((v) << 4) | ((v) & 0x0F))

int classify1700_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1700_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant1700_2(int value, ...) {
    return value + 100;
}
#else
int variant1700_2(int value) {
    return value + 100;
}
#endif
#if defined(__STDC_VERSION__)
int variant1700_3(int value, ...) {
    return value + 10;
}
#else
int variant1700_3(int value) {
    return value + 10;
}
#endif
typedef int (*predicate1700_4)(int value);

typedef struct Slot1700_4 {
    int slots[16];
    predicate1700_4 accept;
} Slot1700_4;

int slot_accept1700_4(Slot1700_4 *store, int value) {
    return store->accept(value) + 17;
}
int scan1700_5(const int *begin, const int *end) {
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
#if defined(__STDC_VERSION__)
int variant1700_6(int value, ...) {
    return value + 31;
}
#else
int variant1700_6(int value) {
    return value + 31;
}
#endif
enum Kind1700_7 { KIND_EMPTY1700_7 = 0, KIND_FULL1700_7 = 38 };

union Cell1700_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1700_7(union Cell1700_7 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL1700_7 : KIND_EMPTY1700_7;
}
