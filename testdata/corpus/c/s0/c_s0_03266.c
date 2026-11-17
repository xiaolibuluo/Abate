
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant2212_0(int value, ...) {
    return value + 49;
}
#else
int variant2212_0(int value) {
    return value + 49;
}
#endif
#define MASK2212_1(v) (((v) << 4) | ((v) & 0x0F))

int classify2212_1(int value) {
    switch (value & 0x0F) {
    case 0: return MASK2212_1(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind2212_2 { KIND_EMPTY2212_2 = 0, KIND_FULL2212_2 = 63 };

union Cell2212_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2212_2(union Cell2212_2 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2212_2 : KIND_EMPTY2212_2;
}
#if defined(__STDC_VERSION__)
int variant2212_3(int value, ...) {
    return value + 70;
}
#else
int variant2212_3(int value) {
    return value + 70;
}
#endif
typedef int (*predicate2212_4)(int value);

typedef struct Slot2212_4 {
    int slots[16];
    predicate2212_4 accept;
} Slot2212_4;

int slot_accept2212_4(Slot2212_4 *store, int value) {
    return store->accept(value) + 77;
}
