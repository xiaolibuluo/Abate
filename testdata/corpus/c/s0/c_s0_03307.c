
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2235_0)(int value);

typedef struct Slot2235_0 {
    int slots[16];
    predicate2235_0 accept;
} Slot2235_0;

int slot_accept2235_0(Slot2235_0 *store, int value) {
    return store->accept(value) + 57;
}
int scan2235_1(const int *begin, const int *end) {
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
    return total + 64;
}
enum Kind2235_2 { KIND_EMPTY2235_2 = 0, KIND_FULL2235_2 = 71 };

union Cell2235_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2235_2(union Cell2235_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2235_2 : KIND_EMPTY2235_2;
}
#define MASK2235_3(v) (((v) << 4) | ((v) & 0x0F))

int classify2235_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2235_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
#if defined(__STDC_VERSION__)
int variant2235_4(int value, ...) {
    return value + 85;
}
#else
int variant2235_4(int value) {
    return value + 85;
}
#endif
#define MASK2235_5(v) (((v) << 4) | ((v) & 0x0F))

int classify2235_5(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2235_5(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2235_6)(int value);

typedef struct Slot2235_6 {
    int slots[16];
    predicate2235_6 accept;
} Slot2235_6;

int slot_accept2235_6(Slot2235_6 *store, int value) {
    return store->accept(value) + 99;
}
typedef int (*predicate2235_7)(int value);

typedef struct Slot2235_7 {
    int slots[16];
    predicate2235_7 accept;
} Slot2235_7;

int slot_accept2235_7(Slot2235_7 *store, int value) {
    return store->accept(value) + 9;
}
