
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1553_0)(int value);

typedef struct Slot1553_0 {
    int slots[16];
    predicate1553_0 accept;
} Slot1553_0;

int slot_accept1553_0(Slot1553_0 *store, int value) {
    return store->accept(value) + 18;
}
#if defined(__STDC_VERSION__)
int variant1553_1(int value, ...) {
    return value + 25;
}
#else
int variant1553_1(int value) {
    return value + 25;
}
#endif
#if defined(__STDC_VERSION__)
int variant1553_2(int value, ...) {
    return value + 32;
}
#else
int variant1553_2(int value) {
    return value + 32;
}
#endif
#define MASK1553_3(v) (((v) << 4) | ((v) & 0x0F))

int classify1553_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1553_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind1553_4 { KIND_EMPTY1553_4 = 0, KIND_FULL1553_4 = 46 };

union Cell1553_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1553_4(union Cell1553_4 *cell) {
    cell->number = 66;
    return cell->number > 0 ? KIND_FULL1553_4 : KIND_EMPTY1553_4;
}
enum Kind1553_5 { KIND_EMPTY1553_5 = 0, KIND_FULL1553_5 = 53 };

union Cell1553_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1553_5(union Cell1553_5 *cell) {
    cell->number = 16;
    return cell->number > 0 ? KIND_FULL1553_5 : KIND_EMPTY1553_5;
}
int scan1553_6(const int *begin, const int *end) {
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
    return total + 60;
}
#define MASK1553_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1553_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1553_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
