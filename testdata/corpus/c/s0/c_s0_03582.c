
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

typedef int (*predicate2418_0)(int value);

typedef struct Slot2418_0 {
    int slots[16];
    predicate2418_0 accept;
} Slot2418_0;

int slot_accept2418_0(Slot2418_0 *store, int value) {
    return store->accept(value) + 11;
}
typedef int (*predicate2418_1)(int value);

typedef struct Slot2418_1 {
    int slots[16];
    predicate2418_1 accept;
} Slot2418_1;

int slot_accept2418_1(Slot2418_1 *store, int value) {
    return store->accept(value) + 18;
}
enum Kind2418_2 { KIND_EMPTY2418_2 = 0, KIND_FULL2418_2 = 25 };

union Cell2418_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2418_2(union Cell2418_2 *cell) {
    cell->number = 60;
    return cell->number > 0 ? KIND_FULL2418_2 : KIND_EMPTY2418_2;
}
#if defined(__STDC_VERSION__)
int variant2418_3(int value, ...) {
    return value + 32;
}
#else
int variant2418_3(int value) {
    return value + 32;
}
#endif
typedef int (*predicate2418_4)(int value);

typedef struct Slot2418_4 {
    int slots[16];
    predicate2418_4 accept;
} Slot2418_4;

int slot_accept2418_4(Slot2418_4 *store, int value) {
    return store->accept(value) + 39;
}
enum Kind2418_5 { KIND_EMPTY2418_5 = 0, KIND_FULL2418_5 = 46 };

union Cell2418_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2418_5(union Cell2418_5 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL2418_5 : KIND_EMPTY2418_5;
}
