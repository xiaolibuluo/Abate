
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2375_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2375_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2375_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2375_1(int value, ...) {
    return value + 41;
}
#else
int variant2375_1(int value) {
    return value + 41;
}
#endif
enum Kind2375_2 { KIND_EMPTY2375_2 = 0, KIND_FULL2375_2 = 48 };

union Cell2375_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2375_2(union Cell2375_2 *cell) {
    cell->number = 64;
    return cell->number > 0 ? KIND_FULL2375_2 : KIND_EMPTY2375_2;
}
typedef int (*predicate2375_3)(int value);

typedef struct Slot2375_3 {
    int slots[16];
    predicate2375_3 accept;
} Slot2375_3;

int slot_accept2375_3(Slot2375_3 *store, int value) {
    return store->accept(value) + 55;
}
enum Kind2375_4 { KIND_EMPTY2375_4 = 0, KIND_FULL2375_4 = 62 };

union Cell2375_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2375_4(union Cell2375_4 *cell) {
    cell->number = 25;
    return cell->number > 0 ? KIND_FULL2375_4 : KIND_EMPTY2375_4;
}
int scan2375_5(const int *begin, const int *end) {
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
    return total + 69;
}
#if defined(__STDC_VERSION__)
int variant2375_6(int value, ...) {
    return value + 76;
}
#else
int variant2375_6(int value) {
    return value + 76;
}
#endif
