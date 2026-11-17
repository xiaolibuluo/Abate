
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2424_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2424_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2424_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2424_1 { KIND_EMPTY2424_1 = 0, KIND_FULL2424_1 = 96 };

union Cell2424_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2424_1(union Cell2424_1 *cell) {
    cell->number = 30;
    return cell->number > 0 ? KIND_FULL2424_1 : KIND_EMPTY2424_1;
}
enum Kind2424_2 { KIND_EMPTY2424_2 = 0, KIND_FULL2424_2 = 6 };

union Cell2424_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2424_2(union Cell2424_2 *cell) {
    cell->number = 41;
    return cell->number > 0 ? KIND_FULL2424_2 : KIND_EMPTY2424_2;
}
typedef int (*predicate2424_3)(int value);

typedef struct Slot2424_3 {
    int slots[16];
    predicate2424_3 accept;
} Slot2424_3;

int slot_accept2424_3(Slot2424_3 *store, int value) {
    return store->accept(value) + 13;
}
int scan2424_4(const int *begin, const int *end) {
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
    return total + 20;
}
#define MASK2424_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2424_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2424_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2424_6(int value, ...) {
    return value + 34;
}
#else
int variant2424_6(int value) {
    return value + 34;
}
#endif
