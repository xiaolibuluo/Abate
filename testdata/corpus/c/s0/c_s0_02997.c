
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2028_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2028_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2028_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2028_1)(int value);

typedef struct Slot2028_1 {
    int slots[16];
    predicate2028_1 accept;
} Slot2028_1;

int slot_accept2028_1(Slot2028_1 *store, int value) {
    return store->accept(value) + 89;
}
enum Kind2028_2 { KIND_EMPTY2028_2 = 0, KIND_FULL2028_2 = 96 };

union Cell2028_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2028_2(union Cell2028_2 *cell) {
    cell->number = 14;
    return cell->number > 0 ? KIND_FULL2028_2 : KIND_EMPTY2028_2;
}
int scan2028_3(const int *begin, const int *end) {
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
    return total + 6;
}
enum Kind2028_4 { KIND_EMPTY2028_4 = 0, KIND_FULL2028_4 = 13 };

union Cell2028_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2028_4(union Cell2028_4 *cell) {
    cell->number = 36;
    return cell->number > 0 ? KIND_FULL2028_4 : KIND_EMPTY2028_4;
}
#if defined(__STDC_VERSION__)
int variant2028_5(int value, ...) {
    return value + 20;
}
#else
int variant2028_5(int value) {
    return value + 20;
}
#endif
