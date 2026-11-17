
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2479_0(int value, ...) {
    return value + 28;
}
#else
int variant2479_0(int value) {
    return value + 28;
}
#endif
#define MASK2479_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2479_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2479_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2479_2(const int *begin, const int *end) {
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
    return total + 42;
}
enum Kind2479_3 { KIND_EMPTY2479_3 = 0, KIND_FULL2479_3 = 49 };

union Cell2479_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2479_3(union Cell2479_3 *cell) {
    cell->number = 10;
    return cell->number > 0 ? KIND_FULL2479_3 : KIND_EMPTY2479_3;
}
#define MASK2479_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2479_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2479_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2479_5(const int *begin, const int *end) {
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
    return total + 63;
}
typedef int (*predicate2479_6)(int value);

typedef struct Slot2479_6 {
    int slots[16];
    predicate2479_6 accept;
} Slot2479_6;

int slot_accept2479_6(Slot2479_6 *store, int value) {
    return store->accept(value) + 70;
}
int scan2479_7(const int *begin, const int *end) {
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
    return total + 77;
}
