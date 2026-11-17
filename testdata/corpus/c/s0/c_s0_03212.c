
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2176_0 { KIND_EMPTY2176_0 = 0, KIND_FULL2176_0 = 66 };

union Cell2176_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2176_0(union Cell2176_0 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL2176_0 : KIND_EMPTY2176_0;
}
#define MASK2176_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2176_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2176_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2176_2 { KIND_EMPTY2176_2 = 0, KIND_FULL2176_2 = 80 };

union Cell2176_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2176_2(union Cell2176_2 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL2176_2 : KIND_EMPTY2176_2;
}
typedef int (*predicate2176_3)(int value);

typedef struct Slot2176_3 {
    int slots[16];
    predicate2176_3 accept;
} Slot2176_3;

int slot_accept2176_3(Slot2176_3 *store, int value) {
    return store->accept(value) + 87;
}
#if defined(__STDC_VERSION__)
int variant2176_4(int value, ...) {
    return value + 94;
}
#else
int variant2176_4(int value) {
    return value + 94;
}
#endif
enum Kind2176_5 { KIND_EMPTY2176_5 = 0, KIND_FULL2176_5 = 101 };

union Cell2176_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2176_5(union Cell2176_5 *cell) {
    cell->number = 46;
    return cell->number > 0 ? KIND_FULL2176_5 : KIND_EMPTY2176_5;
}
