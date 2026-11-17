
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2129_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2129_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2129_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2129_1 { KIND_EMPTY2129_1 = 0, KIND_FULL2129_1 = 44 };

union Cell2129_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2129_1(union Cell2129_1 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2129_1 : KIND_EMPTY2129_1;
}
int scan2129_2(const int *begin, const int *end) {
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
    return total + 51;
}
#if defined(__STDC_VERSION__)
int variant2129_3(int value, ...) {
    return value + 58;
}
#else
int variant2129_3(int value) {
    return value + 58;
}
#endif
enum Kind2129_4 { KIND_EMPTY2129_4 = 0, KIND_FULL2129_4 = 65 };

union Cell2129_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2129_4(union Cell2129_4 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL2129_4 : KIND_EMPTY2129_4;
}
typedef int (*predicate2129_5)(int value);

typedef struct Slot2129_5 {
    int slots[16];
    predicate2129_5 accept;
} Slot2129_5;

int slot_accept2129_5(Slot2129_5 *store, int value) {
    return store->accept(value) + 72;
}
enum Kind2129_6 { KIND_EMPTY2129_6 = 0, KIND_FULL2129_6 = 79 };

union Cell2129_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2129_6(union Cell2129_6 *cell) {
    cell->number = 33;
    return cell->number > 0 ? KIND_FULL2129_6 : KIND_EMPTY2129_6;
}
