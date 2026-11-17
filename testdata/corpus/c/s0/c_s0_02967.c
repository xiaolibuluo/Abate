
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2007_0)(int value);

typedef struct Slot2007_0 {
    int slots[16];
    predicate2007_0 accept;
} Slot2007_0;

int slot_accept2007_0(Slot2007_0 *store, int value) {
    return store->accept(value) + 100;
}
#if defined(__STDC_VERSION__)
int variant2007_1(int value, ...) {
    return value + 10;
}
#else
int variant2007_1(int value) {
    return value + 10;
}
#endif
enum Kind2007_2 { KIND_EMPTY2007_2 = 0, KIND_FULL2007_2 = 17 };

union Cell2007_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2007_2(union Cell2007_2 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL2007_2 : KIND_EMPTY2007_2;
}
