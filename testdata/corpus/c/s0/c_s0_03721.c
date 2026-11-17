
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2517_0 { KIND_EMPTY2517_0 = 0, KIND_FULL2517_0 = 37 };

union Cell2517_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2517_0(union Cell2517_0 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2517_0 : KIND_EMPTY2517_0;
}
#if defined(__STDC_VERSION__)
int variant2517_1(int value, ...) {
    return value + 44;
}
#else
int variant2517_1(int value) {
    return value + 44;
}
#endif
#define MASK2517_2(v) (((v) << 4) | ((v) & 0x0F))

int classify2517_2(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2517_2(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2517_3)(int value);

typedef struct Slot2517_3 {
    int slots[16];
    predicate2517_3 accept;
} Slot2517_3;

int slot_accept2517_3(Slot2517_3 *store, int value) {
    return store->accept(value) + 58;
}
enum Kind2517_4 { KIND_EMPTY2517_4 = 0, KIND_FULL2517_4 = 65 };

union Cell2517_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2517_4(union Cell2517_4 *cell) {
    cell->number = 43;
    return cell->number > 0 ? KIND_FULL2517_4 : KIND_EMPTY2517_4;
}
#if defined(__STDC_VERSION__)
int variant2517_5(int value, ...) {
    return value + 72;
}
#else
int variant2517_5(int value) {
    return value + 72;
}
#endif
typedef int (*predicate2517_6)(int value);

typedef struct Slot2517_6 {
    int slots[16];
    predicate2517_6 accept;
} Slot2517_6;

int slot_accept2517_6(Slot2517_6 *store, int value) {
    return store->accept(value) + 79;
}
