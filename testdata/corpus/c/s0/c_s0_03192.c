
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2161_0(int value, ...) {
    return value + 65;
}
#else
int variant2161_0(int value) {
    return value + 65;
}
#endif
#define MASK2161_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2161_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2161_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
typedef int (*predicate2161_2)(int value);

typedef struct Slot2161_2 {
    int slots[16];
    predicate2161_2 accept;
} Slot2161_2;

int slot_accept2161_2(Slot2161_2 *store, int value) {
    return store->accept(value) + 79;
}
typedef int (*predicate2161_3)(int value);

typedef struct Slot2161_3 {
    int slots[16];
    predicate2161_3 accept;
} Slot2161_3;

int slot_accept2161_3(Slot2161_3 *store, int value) {
    return store->accept(value) + 86;
}
enum Kind2161_4 { KIND_EMPTY2161_4 = 0, KIND_FULL2161_4 = 93 };

union Cell2161_4 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2161_4(union Cell2161_4 *cell) {
    cell->number = 52;
    return cell->number > 0 ? KIND_FULL2161_4 : KIND_EMPTY2161_4;
}
