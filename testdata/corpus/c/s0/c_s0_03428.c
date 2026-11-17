
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2316_0 { KIND_EMPTY2316_0 = 0, KIND_FULL2316_0 = 43 };

union Cell2316_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2316_0(union Cell2316_0 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL2316_0 : KIND_EMPTY2316_0;
}
#define MASK2316_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2316_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2316_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2316_2)(int value);

typedef struct Slot2316_2 {
    int slots[16];
    predicate2316_2 accept;
} Slot2316_2;

int slot_accept2316_2(Slot2316_2 *store, int value) {
    return store->accept(value) + 57;
}
#if defined(__STDC_VERSION__)
int variant2316_3(int value, ...) {
    return value + 64;
}
#else
int variant2316_3(int value) {
    return value + 64;
}
#endif
enum Kind2316_4 { KIND_EMPTY2316_4 = 0, KIND_FULL2316_4 = 71 };

union Cell2316_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2316_4(union Cell2316_4 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2316_4 : KIND_EMPTY2316_4;
}
typedef int (*predicate2316_5)(int value);

typedef struct Slot2316_5 {
    int slots[16];
    predicate2316_5 accept;
} Slot2316_5;

int slot_accept2316_5(Slot2316_5 *store, int value) {
    return store->accept(value) + 78;
}
