
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2415_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2415_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2415_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan2415_1(const int *begin, const int *end) {
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
    return total + 76;
}
typedef int (*predicate2415_2)(int value);

typedef struct Slot2415_2 {
    int slots[16];
    predicate2415_2 accept;
} Slot2415_2;

int slot_accept2415_2(Slot2415_2 *store, int value) {
    return store->accept(value) + 83;
}
#if defined(__STDC_VERSION__)
int variant2415_3(int value, ...) {
    return value + 90;
}
#else
int variant2415_3(int value) {
    return value + 90;
}
#endif
enum Kind2415_4 { KIND_EMPTY2415_4 = 0, KIND_FULL2415_4 = 97 };

union Cell2415_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2415_4(union Cell2415_4 *cell) {
    cell->number = 61;
    return cell->number > 0 ? KIND_FULL2415_4 : KIND_EMPTY2415_4;
}
#define MASK2415_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2415_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2415_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
