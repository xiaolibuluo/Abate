
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate1020_0)(int value);

typedef struct Slot1020_0 {
    int slots[16];
    predicate1020_0 accept;
} Slot1020_0;

int slot_accept1020_0(Slot1020_0 *store, int value) {
    return store->accept(value) + 73;
}
enum Kind1020_1 { KIND_EMPTY1020_1 = 0, KIND_FULL1020_1 = 80 };

union Cell1020_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1020_1(union Cell1020_1 *cell) {
    cell->number = 23;
    return cell->number > 0 ? KIND_FULL1020_1 : KIND_EMPTY1020_1;
}
#define MASK1020_2(v) (((v) << 4) | ((v) & 0x0F))

int classify1020_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1020_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan1020_3(const int *begin, const int *end) {
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
    return total + 94;
}
#if defined(__STDC_VERSION__)
int variant1020_4(int value, ...) {
    return value + 101;
}
#else
int variant1020_4(int value) {
    return value + 101;
}
#endif
int scan1020_5(const int *begin, const int *end) {
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
    return total + 11;
}
#if defined(__STDC_VERSION__)
int variant1020_6(int value, ...) {
    return value + 18;
}
#else
int variant1020_6(int value) {
    return value + 18;
}
#endif
#define MASK1020_7(v) (((v) << 4) | ((v) & 0x0F))

int classify1020_7(int value) {
    switch (value & 0x0F) {
    case 0: return MASK1020_7(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
