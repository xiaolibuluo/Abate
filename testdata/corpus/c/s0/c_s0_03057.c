
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2063_0 { KIND_EMPTY2063_0 = 0, KIND_FULL2063_0 = 52 };

union Cell2063_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2063_0(union Cell2063_0 *cell) {
    cell->number = 54;
    return cell->number > 0 ? KIND_FULL2063_0 : KIND_EMPTY2063_0;
}
typedef int (*predicate2063_1)(int value);

typedef struct Slot2063_1 {
    int slots[16];
    predicate2063_1 accept;
} Slot2063_1;

int slot_accept2063_1(Slot2063_1 *store, int value) {
    return store->accept(value) + 59;
}
typedef int (*predicate2063_2)(int value);

typedef struct Slot2063_2 {
    int slots[16];
    predicate2063_2 accept;
} Slot2063_2;

int slot_accept2063_2(Slot2063_2 *store, int value) {
    return store->accept(value) + 66;
}
#if defined(__STDC_VERSION__)
int variant2063_3(int value, ...) {
    return value + 73;
}
#else
int variant2063_3(int value) {
    return value + 73;
}
#endif
