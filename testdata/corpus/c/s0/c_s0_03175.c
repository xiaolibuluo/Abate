
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2151_0)(int value);

typedef struct Slot2151_0 {
    int slots[16];
    predicate2151_0 accept;
} Slot2151_0;

int slot_accept2151_0(Slot2151_0 *store, int value) {
    return store->accept(value) + 32;
}
#if defined(__STDC_VERSION__)
int variant2151_1(int value, ...) {
    return value + 39;
}
#else
int variant2151_1(int value) {
    return value + 39;
}
#endif
enum Kind2151_2 { KIND_EMPTY2151_2 = 0, KIND_FULL2151_2 = 46 };

union Cell2151_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2151_2(union Cell2151_2 *cell) {
    cell->number = 21;
    return cell->number > 0 ? KIND_FULL2151_2 : KIND_EMPTY2151_2;
}
