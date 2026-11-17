
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2458_0)(int value);

typedef struct Slot2458_0 {
    int slots[16];
    predicate2458_0 accept;
} Slot2458_0;

int slot_accept2458_0(Slot2458_0 *store, int value) {
    return store->accept(value) + 46;
}
enum Kind2458_1 { KIND_EMPTY2458_1 = 0, KIND_FULL2458_1 = 53 };

union Cell2458_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2458_1(union Cell2458_1 *cell) {
    cell->number = 24;
    return cell->number > 0 ? KIND_FULL2458_1 : KIND_EMPTY2458_1;
}
enum Kind2458_2 { KIND_EMPTY2458_2 = 0, KIND_FULL2458_2 = 60 };

union Cell2458_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2458_2(union Cell2458_2 *cell) {
    cell->number = 35;
    return cell->number > 0 ? KIND_FULL2458_2 : KIND_EMPTY2458_2;
}
typedef int (*predicate2458_3)(int value);

typedef struct Slot2458_3 {
    int slots[16];
    predicate2458_3 accept;
} Slot2458_3;

int slot_accept2458_3(Slot2458_3 *store, int value) {
    return store->accept(value) + 67;
}
#if defined(__STDC_VERSION__)
int variant2458_4(int value, ...) {
    return value + 74;
}
#else
int variant2458_4(int value) {
    return value + 74;
}
#endif
