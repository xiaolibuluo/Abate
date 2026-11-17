
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2640_0)(int value);

typedef struct Slot2640_0 {
    int slots[16];
    predicate2640_0 accept;
} Slot2640_0;

int slot_accept2640_0(Slot2640_0 *store, int value) {
    return store->accept(value) + 84;
}
#if defined(__STDC_VERSION__)
int variant2640_1(int value, ...) {
    return value + 91;
}
#else
int variant2640_1(int value) {
    return value + 91;
}
#endif
enum Kind2640_2 { KIND_EMPTY2640_2 = 0, KIND_FULL2640_2 = 98 };

union Cell2640_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2640_2(union Cell2640_2 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2640_2 : KIND_EMPTY2640_2;
}
