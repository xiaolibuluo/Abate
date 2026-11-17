
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#define MASK2043_0(v) (((v) << 4) | ((v) & 0x0F))

int classify2043_0(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2043_0(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2043_1 { KIND_EMPTY2043_1 = 0, KIND_FULL2043_1 = 90 };

union Cell2043_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2043_1(union Cell2043_1 *cell) {
    cell->number = 47;
    return cell->number > 0 ? KIND_FULL2043_1 : KIND_EMPTY2043_1;
}
#if defined(__STDC_VERSION__)
int variant2043_2(int value, ...) {
    return value + 97;
}
#else
int variant2043_2(int value) {
    return value + 97;
}
#endif
enum Kind2043_3 { KIND_EMPTY2043_3 = 0, KIND_FULL2043_3 = 7 };

union Cell2043_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2043_3(union Cell2043_3 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL2043_3 : KIND_EMPTY2043_3;
}
#define MASK2043_4(v) (((v) << 4) | ((v) & 0x0F))

int classify2043_4(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2043_4(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2043_5)(int value);

typedef struct Slot2043_5 {
    int slots[16];
    predicate2043_5 accept;
} Slot2043_5;

int slot_accept2043_5(Slot2043_5 *store, int value) {
    return store->accept(value) + 21;
}
int scan2043_6(const int *begin, const int *end) {
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
    return total + 28;
}
typedef int (*predicate2043_7)(int value);

typedef struct Slot2043_7 {
    int slots[16];
    predicate2043_7 accept;
} Slot2043_7;

int slot_accept2043_7(Slot2043_7 *store, int value) {
    return store->accept(value) + 35;
}
