
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2515_0)(int value);

typedef struct Slot2515_0 {
    int slots[16];
    predicate2515_0 accept;
} Slot2515_0;

int slot_accept2515_0(Slot2515_0 *store, int value) {
    return store->accept(value) + 11;
}
#if defined(__STDC_VERSION__)
int variant2515_1(int value, ...) {
    return value + 18;
}
#else
int variant2515_1(int value) {
    return value + 18;
}
#endif
#define MASK2515_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2515_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2515_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2515_3 { KIND_EMPTY2515_3 = 0, KIND_FULL2515_3 = 32 };

union Cell2515_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2515_3(union Cell2515_3 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL2515_3 : KIND_EMPTY2515_3;
}
#define MASK2515_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2515_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2515_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2515_5(int value, ...) {
    return value + 46;
}
#else
int variant2515_5(int value) {
    return value + 46;
}
#endif
int scan2515_6(const int *begin, const int *end) {
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
    return total + 53;
}
